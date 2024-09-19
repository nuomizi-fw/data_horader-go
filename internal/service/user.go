package service

import "github.com/nuomizi-fw/data_horader-go/internal/core"

type UserService interface {
	GetUser() error
	GetUsers() error
	CreateUser() error
	UpdateUser() error
	DeleteUser() error
	SetUserRole() error
	ResetPassword() error
	RefreshToken() error
}

type userService struct {
	db     core.DataHoraderDB
	logger core.DataHoraderLogger
}

func NewUserService(
	db core.DataHoraderDB,
	logger core.DataHoraderLogger,
) UserService {
	return &userService{db, logger}
}

func (us *userService) GetUser() error {
	return nil
}

func (us *userService) GetUsers() error {
	return nil
}

func (us *userService) CreateUser() error {
	return nil
}

func (us *userService) UpdateUser() error {
	return nil
}

func (us *userService) DeleteUser() error {
	return nil
}

func (us *userService) SetUserRole() error {
	return nil
}

func (us *userService) ResetPassword() error {
	return nil
}

func (ur *userService) RefreshToken() error {
	return nil
}
