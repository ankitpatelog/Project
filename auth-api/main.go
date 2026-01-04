package main

import (
	"auth-api/config"
	"auth-api/controllers"
	"auth-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/profile")
	auth.Use(middleware.AuthMiddleware(config.GetJWTSecret()))
	{
		auth.GET("",controllers.Profile)
	}

	r.Run(":8080")

}