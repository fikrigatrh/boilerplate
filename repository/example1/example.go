package example1

import (
	"boilerplate/repository"
	"gorm.io/gorm"
)

type Example1RepoStrtuct struct {
	db *gorm.DB
}

func NewExample1RepoImpl() repository.Example1RepoInterface {
	return &Example1RepoStrtuct{}
}

func (e Example1RepoStrtuct ) name()  {

}