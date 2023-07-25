package controller

import (
	"net/http"
	"online_food_market/helper"
	"online_food_market/model"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var input model.RegisterInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "not enough input!"})
		return
	}

	if input.Password != input.Password2 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "passwords are not same!"})
		return
	}

	user := model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func Login(context *gin.Context) {
	var input model.LoginInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Incomplete JSON input. Please provide all required fields"})
		return
	}

	user, err := model.FindUserByEmail(input.Email)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "user does not exists"})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "password or email does not exists"})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
