package main

import (
	"auth/controllers"
	"auth/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()
	router.POST("/signup", controllers.SignUp)

	
	router.Run() // listens on 0.0.0.0:8080 by default
}