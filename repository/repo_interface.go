package repository

import "boilerplate/models"

type Example1RepoInterface interface {
	AddRole(rl models.Role) (models.Role, error)
	UpdateRole(rl models.Role) (models.Role, error)

}

type LoginRepoInterface interface {
	GetAdminId(name string) (models.Role, error)
	CreateAuth(username string, roleId string) (*models.Auth, error)
}