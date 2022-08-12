package main

import (
	"net/http"
	"screechr/src/controller"
	"screechr/src/dto"
	"screechr/src/middleware"
	"screechr/src/service"

	"github.com/gin-gonic/gin"
)

var loginService service.LoginService = service.StaticLoginService()
var jwtService service.JWTService = service.JWTAuthService()
var profileService service.ProfileService = service.Profile()
var screechService service.ScreechService = service.Screech()
var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

func main() {
	Engine().Run()
}

func Engine() *gin.Engine {
	server := gin.New()
	server.Use(gin.Recovery())
	server.POST("/login", Login)

	operator := server.Group("/")

	operator.Use(middleware.AuthorizeOperator())
	{
		operator.GET("/get_profile", GetProfile)
		operator.GET("/update_profile", UpdateProfile)
		operator.GET("/update_profile_picture", UpdatePicture)
		operator.GET("/get_screeches", GetScreeches)
		operator.GET("/get_screech", GetScreech)
		operator.GET("/create_screech", CreateScreech)
		operator.GET("/update_screech", UpdateScreech)
	}

	authorized := server.Group("/admin")

	authorized.Use(middleware.AuthorizeAdmin())
	{
		authorized.GET("/get_profile", GetProfileByID)
	}
	return server
}

// Login request username and password return access token
func Login(ctx *gin.Context) {
	token := loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, "Unauthorized")
	}
}

// GetProfileByID retrieve profile by id -- input:json {"id": $id}
func GetProfileByID(ctx *gin.Context) {
	data := &dto.ProfileRequestType{}
	ctx.BindJSON(data)
	if data.ID == 0 {
		ctx.JSON(http.StatusBadRequest, "need id")
		return
	}
	profile, err := profileService.GetProfile(data.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Unable to find profile")
	} else {
		ctx.JSON(http.StatusOK, profile)
	}
}

// GetProfile retrieve self profile -- input:json {"id": $id, "token": $token}
func GetProfile(ctx *gin.Context) {
	data := &dto.ProfileRequestType{}
	ctx.BindJSON(data)
	if data.Token == "" {
		ctx.JSON(http.StatusBadRequest, "need token")
		return
	}
	if data.ID == 0 {
		ctx.JSON(http.StatusBadRequest, "need id")
		return
	}
	profile, err := profileService.GetProfile(data.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Unable to find profile")
		return
	}
	if data.Token != profile.Token {
		ctx.JSON(http.StatusBadRequest, "invalid token")
		return
	}
	ctx.JSON(http.StatusOK, profile)
}

// UpdateProfile update entire profile -- input:json {"id": $id,"user_name": $uid, "first_name": $name, "last_name": $name, "url": $url, "token": $token}
func UpdateProfile(ctx *gin.Context) {
	data := &dto.ProfileRequestType{}
	ctx.BindJSON(data)
	if data.Token == "" {
		ctx.JSON(http.StatusBadRequest, "need token")
		return
	}
	if data.ID == 0 {
		ctx.JSON(http.StatusBadRequest, "need id")
		return
	}
	err := profileService.UpdateProfile(data.ID, data.UserName, data.FirstName, data.LastName, data.Token, data.ImageUrl)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Unable to find profile")
		return
	}
	ctx.JSON(http.StatusOK, "profile updated")
}

// UpdatePicture update picture in profile -- input:json {"id": $id, "url": $url}
func UpdatePicture(ctx *gin.Context) {
	data := &dto.ProfileRequestType{}
	ctx.BindJSON(data)
	if data.Token == "" {
		ctx.JSON(http.StatusBadRequest, "need token")
		return
	}
	if data.ID == 0 {
		ctx.JSON(http.StatusBadRequest, "need id")
		return
	}
	if data.ImageUrl == "" {
		ctx.JSON(http.StatusBadRequest, "need image url")
		return
	}
	err := profileService.UpdatePicture(data.ID, data.ImageUrl)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Unable to find profile")
		return
	}
	ctx.JSON(http.StatusOK, "picture updated")
}

// GetScreeches retrieve all screeches (sort by OrderByAscend and return specific user's if uid is not 0) -- input:json {"order_by_ascend": $false, "user_id": $id}
func GetScreeches(ctx *gin.Context) {
	data := &dto.ScreechRequestType{}
	ctx.BindJSON(data)
	screeches := screechService.GetScreeches(data.OrderByAscend, data.CreatorID)
	ctx.JSON(http.StatusOK, screeches)
}

// GetScreech retrieve a screech by id -- input:json {"id":$id}
func GetScreech(ctx *gin.Context) {
	data := &dto.ScreechRequestType{}
	ctx.BindJSON(data)
	if data.ID == 0 {
		ctx.JSON(http.StatusBadRequest, "need token")
		return
	}
	screech, err := screechService.GetScreech(data.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Unable to find screech")
		return
	}
	ctx.JSON(http.StatusOK, screech)
}

// GetScreech create a screech -- input:json {"user_id":$id, "content":$content}
func CreateScreech(ctx *gin.Context) {
	data := &dto.ScreechRequestType{}
	ctx.BindJSON(data)
	if data.CreatorID == 0 {
		ctx.JSON(http.StatusBadRequest, "need creator id")
		return
	}
	err := screechService.CreateScreech(data.CreatorID, data.Content)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Unable to create screech")
		return
	}
	ctx.JSON(http.StatusOK, "createScreech")
}

// UpdateScreech update a screech  -- input:json {"id":$id, "content":$content}
func UpdateScreech(ctx *gin.Context) {
	data := &dto.ScreechRequestType{}
	ctx.BindJSON(data)
	if data.ID == 0 {
		ctx.JSON(http.StatusBadRequest, "need id")
		return
	}
	err := screechService.UpdateContent(data.ID, data.Content)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Unable to create screech")
		return
	}
	ctx.JSON(http.StatusOK, "createScreech")
}
