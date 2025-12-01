package main

import (
	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/routes"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	config.DBconnect()
	config.SyncDatabase()
}

func main() {
	routes.UserRouter(r)
	r.Run(":8080")
}
