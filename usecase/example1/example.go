package example1

import (
	"boilerplate/config/log"
	"boilerplate/models"
	"boilerplate/models/contract"
	"boilerplate/repository"
	"boilerplate/usecase"
	"errors"
)

type Example1UsecaseStrtuct struct {
	ex repository.Example1RepoInterface
	log *log.LogCustom
}

func NewUserUsecaseImpl(ex repository.Example1RepoInterface, log *log.LogCustom) usecase.Example1UsecaseInterface {
	return &Example1UsecaseStrtuct{ex, log}
}

func (e Example1UsecaseStrtuct ) AddRole(rl models.Role) (models.Role, error) {
	if rl.RoleName != "accounting" {
		if  rl.RoleName != "general_support" {
			if rl.RoleName != "admin" {
				if rl.RoleName != "customer" {
					e.log.Error(errors.New("invalid param"), "usecase: add role", nil)
					return models.Role{}, errors.New(contract.ErrInvalidRoleName)
				}
			}
		}
	}

	role, err := e.ex.AddRole(rl)
	if err != nil {
		return models.Role{}, err
	}

	return role,  nil
}

func (e Example1UsecaseStrtuct) UpdateRole(rl models.Role) (models.Role, error) {
	role, err := e.ex.UpdateRole(rl)
	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}