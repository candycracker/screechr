package controller

import (
	"screechr/src/dto"
	"screechr/src/service"

	"github.com/gin-gonic/gin"
)

// login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func LoginHandler(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential dto.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	role, uid, isUserAuthenticated := controller.loginService.LoginUser(credential.UserName, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.UserName, credential.Password, uid, role)

	}
	return ""
}
