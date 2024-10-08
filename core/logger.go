package core

import (
	"context"
	"os"
	"time"

	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	gormlogger "gorm.io/gorm/logger"
)

var zapLogger *zap.Logger

type DataHoraderLogger struct {
	*zap.SugaredLogger
}

type GormLogger struct {
	*DataHoraderLogger
	gormlogger.Config
}

type FxLogger struct {
	*DataHoraderLogger
}

type FiberLogger struct {
	*DataHoraderLogger
}

func NewDataHoraderLogger() DataHoraderLogger {
	return newDataHoraderLogger(NewDataHoraderConfig())
}

func newDataHoraderLogger(config DataHoraderConfig) DataHoraderLogger {
	var zapConfig zap.Config

	if _, err := os.Stat(config.Logger.LogPath); os.IsNotExist(err) {
		if err := os.MkdirAll(config.Logger.LogPath, 0o750); err != nil {
			panic(err)
		}
	}

	logOutputPath := config.Logger.LogPath + "/" + config.Logger.LogName + "." + config.Logger.LogExt

	if config.Server.Debug {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		zapConfig = zap.NewProductionConfig()
		zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		zapConfig.OutputPaths = []string{logOutputPath}
	}

	switch config.Logger.LogLevel {
	case "debug":
		zapConfig.Level.SetLevel(zapcore.DebugLevel)
	case "info":
		zapConfig.Level.SetLevel(zapcore.InfoLevel)
	case "warn":
		zapConfig.Level.SetLevel(zapcore.WarnLevel)
	case "error":
		zapConfig.Level.SetLevel(zapcore.ErrorLevel)
	case "fatal":
		zapConfig.Level.SetLevel(zapcore.FatalLevel)
	default:
		zapConfig.Level.SetLevel(zapcore.PanicLevel)
	}

	zapLogger, _ = zapConfig.Build()

	return *newDataHoraderSugaredLogger(zapLogger)
}

func newDataHoraderSugaredLogger(logger *zap.Logger) *DataHoraderLogger {
	return &DataHoraderLogger{
		SugaredLogger: logger.Sugar(),
	}
}

func (l *DataHoraderLogger) GetGormLogger() gormlogger.Interface {
	logger := zapLogger.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(3),
	)

	return &GormLogger{
		DataHoraderLogger: newDataHoraderSugaredLogger(logger),
		Config: gormlogger.Config{
			LogLevel: gormlogger.Info,
		},
	}
}

// GetFxLogger gets logger for go-fx
func (l *DataHoraderLogger) GetFxLogger() fxevent.Logger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)
	return &FxLogger{newDataHoraderSugaredLogger(logger)}
}

func (l *FxLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.DataHoraderLogger.Debug("OnStart hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.DataHoraderLogger.Debug("OnStart hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			l.DataHoraderLogger.Debug("OnStart hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.OnStopExecuting:
		l.DataHoraderLogger.Debug("OnStop hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.DataHoraderLogger.Debug("OnStop hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			l.DataHoraderLogger.Debug("OnStop hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.Supplied:
		l.DataHoraderLogger.Debug("supplied: ", zap.String("type", e.TypeName), zap.Error(e.Err))
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.DataHoraderLogger.Debug("provided: ", e.ConstructorName, " => ", rtype)
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.DataHoraderLogger.Debug("decorated: ",
				zap.String("decorator", e.DecoratorName),
				zap.String("type", rtype),
			)
		}
	case *fxevent.Invoking:
		l.DataHoraderLogger.Debug("invoking: ", e.FunctionName)
	case *fxevent.Started:
		if e.Err == nil {
			l.DataHoraderLogger.Debug("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err == nil {
			l.DataHoraderLogger.Debug("initialized: custom fxevent.Logger -> ", e.ConstructorName)
		}
	}
}

// ------ GORM logger interface implementation -----

// LogMode set log mode
func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info prints info
func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Info {
		l.Debugf(str, args...)
	}
}

// Warn prints warn messages
func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Warn {
		l.Warnf(str, args...)
	}
}

// Error prints error messages
func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel >= gormlogger.Error {
		l.Errorf(str, args...)
	}
}

// Trace prints trace messages
func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	if l.LogLevel >= gormlogger.Info {
		sql, rows := fc()
		l.Debug("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		return
	}

	if l.LogLevel >= gormlogger.Warn {
		sql, rows := fc()
		l.SugaredLogger.Warn("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		return
	}

	if l.LogLevel >= gormlogger.Error {
		sql, rows := fc()
		l.SugaredLogger.Error("[", elapsed.Milliseconds(), " ms, ", rows, " rows] ", "sql -> ", sql)
		return
	}
}

// Printf prints go-fx logs
func (l FxLogger) Printf(str string, args ...interface{}) {
	if len(args) > 0 {
		l.Debugf(str, args)
	}
	l.Debug(str)
}

// Writer interface implementation for fiber-framework
func (l FiberLogger) Write(p []byte) (n int, err error) {
	l.Info(string(p))
	return len(p), nil
}
