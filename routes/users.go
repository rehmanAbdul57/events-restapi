package routes

import (
	"example.com/RestAPI/models"
	"example.com/RestAPI/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse request data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse request data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Could not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not authenticate user"})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
