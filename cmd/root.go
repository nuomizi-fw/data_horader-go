package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2/log"
	"github.com/nuomizi-fw/data_horader-go/core"
	"github.com/nuomizi-fw/data_horader-go/middleware"
	"github.com/nuomizi-fw/data_horader-go/model"
	"github.com/nuomizi-fw/data_horader-go/router"
	"github.com/nuomizi-fw/data_horader-go/service"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use:   "data_horader",
	Short: "An All-in-One self-hosted solution for your videos, music, manga, novels and more.",
	Long:  "An All-in-One self-hosted solution for your videos, music, manga, novels and more.",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		kill := make(chan os.Signal, 1)
		signal.Notify(kill, syscall.SIGTERM, syscall.SIGINT)

		go func() {
			<-kill
			cancel()
		}()

		app := fx.New(
			fx.WithLogger(func() fxevent.Logger {
				logger := core.NewDataHoraderLogger()
				return logger.GetFxLogger()
			}),
			core.Module,
			middleware.Module,
			model.Module,
			router.Module,
			service.Module,
			fx.Invoke(StartDataHorader),
		)

		if err := app.Start(ctx); err != nil {
			if err != context.Canceled {
				log.Fatalf("Failed to start app: %s", err)
			}
		}

		<-ctx.Done()
	},
}

func StartDataHorader(
	lc fx.Lifecycle,
	config core.DataHoraderConfig,
	db core.DataHoraderDB,
	logger core.DataHoraderLogger,
	server core.DataHoraderServer,
	middleware middleware.DataHoraderMiddlewares,
	model model.DataHoraderModel,
	router router.DataHoraderRouters,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			router.InitRouter()
			middleware.InitMiddleware()
			model.AutoMigrate()

			go func() {
				if config.Server.TLS.Enabled {
					if err := server.App.ListenTLS(config.Server.Port, config.Server.TLS.CertFile, config.Server.TLS.KeyFile); err != nil {
						logger.Panic("Failed to start https server: %s", zap.Error(err))
					}
				} else {
					if err := server.App.Listen(config.Server.Port); err != nil {
						logger.Panic("Failed to start http server: %s", zap.Error(err))
					}
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := server.App.Shutdown(); err != nil {
				return err
			}
			return nil
		},
	})
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
