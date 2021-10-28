package usecase

import "boilerplate/models"

type Example1UsecaseInterface interface {
	AddRole(rl models.Role) (models.Role, error)
	UpdateRole(rl models.Role) (models.Role, error)
}

type Example2UsecaseInterface interface {
	IniFuncSatu()
}

type LoginUsecaseInterface interface {
	GetAdminId(name string) (models.Role, error)
	CreateAuth(username string, roleId string) (*models.Auth, error)
	SignIn(authD models.Auth) (string, error)
}

type ErrorHandlerUsecase interface {
	ResponseError(error interface{}) (int, interface{})
	ValidateRequest(error interface{}) error
}