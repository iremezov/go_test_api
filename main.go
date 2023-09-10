package main

import (
	"awesomeProject/controllers"
	"awesomeProject/helpers"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
)

func main() {

	helpers.Log_msg("Main", "Start application")

	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/users", controllers.FindUsers)
	r.GET("/users/:id", controllers.FindUser)
	r.POST("/users", controllers.CreateUser)
	r.POST("/users/update/:id", controllers.UpdateUser)

	r.Run(":8080")

}
