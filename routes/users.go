package routes

import (
	"net/http"
	"strconv"

	"chaitanyaallu.dev/event-management/models"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request body",
		})
		return
	}
	err = user.CreateUser()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request body",
		})
		return
	}
	err = models.ValidateCredentials(user.Email, user.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

func getUserByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID",
		})
		return
	}
	user, err := models.GetUserByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
