package controller

import (
	"boilerplate/config/env"
	"boilerplate/models"
	"boilerplate/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LoginController struct {
	loginUsecase usecase.LoginUsecaseInterface
}

func NewLoginControllerImpl(r *gin.RouterGroup, loginUsecase usecase.LoginUsecaseInterface) {
	handler := LoginController{loginUsecase}

	r.POST("/login", handler.Login)
}

func (l LoginController) Login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "opps something went wrong")
		return
	}

	if user.Username == env.Config.UsernameSU && user.Password == env.Config.PassSu {
		name := "admin"
		admin, err := l.loginUsecase.GetAdminId(name)
		if err != nil {
			return
		}
		id := int(admin.ID)
		idRes := strconv.Itoa(id)
		user.RoleId = idRes
	}

	var authD models.Auth

	auth, err := l.loginUsecase.CreateAuth(user.Username, user.RoleId)
	if err != nil {
		return 
	}

	authD.AuthUUID = auth.AuthUUID
	authD.Username = auth.Username
	authD.RoleId = auth.RoleId

	in, err := l.loginUsecase.SignIn(authD)
	if err != nil {
		return 
	}

	var JWT models.TokenStruct

	JWT.Token = in

	c.JSON(http.StatusOK, JWT)

}

func (l LoginController) GetAll(c *gin.Context)  {
	//TODO function get all aja yg ada di usecase


}
