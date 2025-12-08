package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"main.go/config"
	"main.go/models"
)

func UserLoginHandler(c *gin.Context) {
	var user models.UserModel
	var UserInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&UserInput); err != nil {
		log.Println("Failed to bind login data")
		c.JSON(400, gin.H{
			"message": "Failed to bind input data",
			"err":     err.Error(),
		})
		return
	}

	if UserInput.Email == "" || UserInput.Password == "" {
		log.Println("Missing email or password")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email and Password are required",
		})
		return
	}

	if err := config.DB.Where("email = ?", UserInput.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found",
			"code":http.StatusUnauthorized,
			"err":     err.Error(),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid password",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successfull",
	})
}