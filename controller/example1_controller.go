package controller

import (
	"boilerplate/config/log"
	"boilerplate/models"
	"boilerplate/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Example1Controller struct {
	ex usecase.Example1UsecaseInterface
	log *log.LogCustom
	err usecase.ErrorHandlerUsecase
}

func NewExampleControllerImpl(r *gin.RouterGroup, ex usecase.Example1UsecaseInterface, log *log.LogCustom, err usecase.ErrorHandlerUsecase) {
	handler := Example1Controller{ex, log, err}

	r.POST("/addRole", handler.AddRole)
	r.PUT("/updateRole", handler.UpdateRole)
}

func (e Example1Controller) AddRole(c *gin.Context)  {
	var rl models.Role

	err := c.ShouldBindJSON(&rl)
	if err != nil {
		e.log.Error(err, "controller: add role", nil)
		c.JSON(http.StatusBadRequest, "opps")
		return
	}

	err = e.err.ValidateRequest(rl)
	if err != nil {
		e.log.Error(err, "controller: validate request data", nil)
		c.Error(err)
		c.Abort()
		return
	}

	role, err := e.ex.AddRole(rl)
	if err != nil {
		e.log.Error(err, "controller: call function usecase add role", nil)
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, role)
}

func (e Example1Controller) UpdateRole(c *gin.Context) {
	var role models.Role

	err := c.ShouldBindJSON(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error when binding")
		return
	}

	roleData, err := e.ex.UpdateRole(role)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ooppsssssss")
		return
	}

	c.JSON(http.StatusOK, roleData)
}