package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nuomizi-fw/data_horader-go/internal/core"
)

type UserRouter struct {
	data_horader core.DataHoraderServer
	logger       core.DataHoraderLogger
}

func NewUserRouter(data_horader core.DataHoraderServer, logger core.DataHoraderLogger) UserRouter {
	return UserRouter{
		data_horader: data_horader,
		logger:       logger,
	}
}

func (ur UserRouter) InitRouter() {
	ur.logger.Info("Initializing user router")
}

func (ur *UserRouter) GetUser(ctx *fiber.Ctx) error {
	return nil
}

func (ur *UserRouter) GetUsers(ctx *fiber.Ctx) error {
	return nil
}

func (ur *UserRouter) CreateUser(ctx *fiber.Ctx) error {
	return nil
}

func (ur *UserRouter) UpdateUser(ctx *fiber.Ctx) error {
	return nil
}

func (ur *UserRouter) DeleteUser(ctx *fiber.Ctx) error {
	return nil
}

func (ur *UserRouter) SetUserRole(ctx *fiber.Ctx) error {
	return nil
}

func (ur *UserRouter) SetUserPermission(ctx *fiber.Ctx) error {
	return nil
}

func (ur *UserRouter) ResetPassword(ctx *fiber.Ctx) error {
	return nil
}

func (ur *UserRouter) RefreshToken(ctx *fiber.Ctx) error {
	return nil
}
