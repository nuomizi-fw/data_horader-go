package service

import "github.com/nuomizi-fw/data_horader-go/internal/core"

type AuthService interface {
	Register() error
	Login() error
	MFALogin() error
	ForgotPassword() error
}

type authService struct {
	db     core.DataHoraderDB
	logger core.DataHoraderLogger
}

func NewAuthService(
	db core.DataHoraderDB,
	logger core.DataHoraderLogger,
) AuthService {
	return &authService{db, logger}
}

func (as *authService) Register() error {
	return nil
}

func (as *authService) Login() error {
	return nil
}

func (as *authService) MFALogin() error {
	return nil
}

func (as *authService) ForgotPassword() error {
	return nil
}
