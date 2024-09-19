package middleware

import "github.com/nuomizi-fw/data_horader-go/internal/core"

type CorsMiddleware struct {
	config       core.DataHoraderConfig
	logger       core.DataHoraderLogger
	data_horader core.DataHoraderServer
}

func NewCorsMiddleware(
	config core.DataHoraderConfig,
	logger core.DataHoraderLogger,
	data_horader core.DataHoraderServer,
) CorsMiddleware {
	return CorsMiddleware{config, logger, data_horader}
}

func (cm CorsMiddleware) InitMiddleware() {
	cm.logger.Info("Initializing Cors middleware")

	if !cm.config.Server.Cors.Enabled {
		cm.logger.Info("Cors middleware disabled")
		return
	}
}
