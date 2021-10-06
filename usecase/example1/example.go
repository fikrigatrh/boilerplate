package example1

import (
	"boilerplate/usecase"
)

type Example1UsecaseStrtuct struct {

}

func NewUserUsecaseImpl() usecase.Example1UsecaseInterface {
	return &Example1UsecaseStrtuct{}
}

func (e Example1UsecaseStrtuct ) name()  {

}