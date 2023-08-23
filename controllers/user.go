package controllers

import (
	"LoginProject_Backend/models"
	"LoginProject_Backend/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserView views.UserView
}

func New(userview views.UserView) UserController {
	return UserController{
		UserView: userview,
	}
}

func (uc *UserController) CheckUser(ctx *gin.Context) {
	email := ctx.Param("email")
	err := uc.UserView.CheckUser(&email)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) MakeUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.UserView.MakeUser(&user.Email, &user.Name, &user.Picture)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	email := ctx.Param("email")
	user, err := uc.UserView.GetUser(&email)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserView.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) LoginUser(ctx *gin.Context) {
	var user *models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// err := uc.UserView.MakeUser(&user.Email, &user.Name, &user.Picture)
	user, err := uc.UserView.LoginUser(&user.Username, &user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"email": user.Email})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.GET("/check/:email", uc.CheckUser)
	userroute.POST("/make", uc.MakeUser)
	userroute.GET("/get/:email", uc.GetUser)
	userroute.PATCH("/update", uc.UpdateUser)
	userroute.POST("/login", uc.LoginUser)
}
