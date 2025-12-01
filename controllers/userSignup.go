package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"main.go/config"
	"main.go/models"
)

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password")
		return "", err
	}
	return string(hashedBytes), nil
}

func SignUp(c *gin.Context) {

	var userInput models.UserModel

	if err := c.ShouldBindJSON(&userInput); err != nil {
		log.Println("Failed to bind signup data")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Binding the data",
			"err":     err.Error(),
		})
		return
	}

	if userInput.Name == "" || userInput.Email == "" || userInput.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "All fields are required",
		})
		return
	}

	hashedPassword, err := HashPassword(userInput.Password)
	if err != nil {
		log.Println("Failed to hash password")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to process password",
			"err":     err.Error(),
		})
		return
	}
	userInput.Password = hashedPassword
	result := config.DB.Create(&userInput)
	if result.Error != nil {
		log.Println("User already exist", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User already exist",
			"err":     result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
