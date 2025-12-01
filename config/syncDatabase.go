package config

import (
	"fmt"
	"main.go/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		fmt.Println("Failed to migrate model")
		return
	}

}