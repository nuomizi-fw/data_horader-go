package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nuomizi-fw/data_horader-go/core"
	"github.com/nuomizi-fw/data_horader-go/service"
)

type AuthRouter struct {
	data_horader core.DataHoraderServer
	logger       core.DataHoraderLogger
	auth         service.AuthService
}

func NewAuthRouter(
	data_horader core.DataHoraderServer,
	logger core.DataHoraderLogger,
	auth service.AuthService,
) AuthRouter {
	return AuthRouter{
		data_horader: data_horader,
		logger:       logger,
		auth:         auth,
	}
}

func (ar AuthRouter) InitRouter() {
	ar.logger.Info("Initializing auth router")

	auth := ar.data_horader.App.Group("/auth")
	{
		auth.Post("/register", ar.Register)
		auth.Post("/login", ar.Login)
		auth.Post("/mfa/generate", ar.MFAGenerate)
		auth.Post("/mfa/login", ar.MFALogin)
		auth.Post("/forgot-password", ar.ForgotPassword)
	}
}

func (as *AuthRouter) Register(ctx *fiber.Ctx) error {
	return nil
}

func (as *AuthRouter) Login(ctx *fiber.Ctx) error {
	return nil
}

func (as *AuthRouter) MFAGenerate(ctx *fiber.Ctx) error {
	return nil
}

func (as *AuthRouter) MFALogin(ctx *fiber.Ctx) error {
	return nil
}

func (as *AuthRouter) ForgotPassword(ctx *fiber.Ctx) error {
	return nil
}
