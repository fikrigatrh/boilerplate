package example2

import (
	"boilerplate/repository"
	"gorm.io/gorm"
)

type Example2RepoStrtuct struct {
	db *gorm.DB
}

func NewExample2RepoImpl() repository.Example2RepoInterface {
	return &Example2RepoStrtuct{}
}

func (e Example2RepoStrtuct ) name()  {

}
