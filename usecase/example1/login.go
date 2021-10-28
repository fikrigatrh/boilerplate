package example1

import (
	"boilerplate/auth"
	"boilerplate/models"
	"boilerplate/repository"
	"boilerplate/usecase"
)

type LoginUsecaseStruct struct {
	loginRepo repository.LoginRepoInterface
}

func NewLoginUsecaseImpl(loginRepo repository.LoginRepoInterface) usecase.LoginUsecaseInterface {
	return &LoginUsecaseStruct{loginRepo}
}

func (l LoginUsecaseStruct) GetAdminId(name string) (models.Role, error)  {
	res, err := l.loginRepo.GetAdminId(name)
	if err != nil {
		return models.Role{}, err
	}

	return res, nil
}

func (l LoginUsecaseStruct) CreateAuth(username string, roleId string) (*models.Auth, error) {
	dataAuth, err := l.loginRepo.CreateAuth(username, roleId)
	if err != nil {
		return nil, err
	}

	return dataAuth, nil
}

func (l LoginUsecaseStruct) SignIn(authD models.Auth) (string, error) {
	token, err := auth.CreateToken(authD)
	if err != nil {
		return "", err
	}
	return token, nil
}