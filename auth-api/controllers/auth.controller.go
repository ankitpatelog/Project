package controllers

import (
	"auth-api/config"
	"auth-api/models"
	"auth-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []models.User{}
var nextID = 1

func Register(c *gin.Context)  {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hash,_ := services.HashPassword(input.Password)

	user := models.User{
		ID:       nextID,
		Name:     input.Name,
		Email:    input.Email,
		Password: hash,
		Role:     "user",
	}

	nextID++
	users = append(users, user)

	c.JSON(201, gin.H{"message": "User registered"})

	//make the input struct that is the input by the user basically its a model
	//then hash the password
	//them user the userModel to store the data ito db or localDb
	//send the response to the user
}

func Login(c *gin.Context)  {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err:=c.ShouldBindJSON(&input)
	if err!=nil {
		c.JSON(http.StatusBadGateway,gin.H{"message":"Input field not correct"})
		return
	}

	//check for the email in the user
	for _, user := range users {
		if user.Email==input.Email {
			//email found
			services.CheckPassword(user.Password,input.Password)

			token, _ := services.GenerateToken(
				user.ID,
				user.Role,
				config.GetJWTSecret(),
			)

			c.JSON(http.StatusOK,gin.H{"token":token})
			return
		}
	}
	c.JSON(401,gin.H{"error":"Invalid credentials"})
}

func Profile(c *gin.Context)  {
	userID := c.GetInt("user_id")
	role := c.GetString("role")

	c.JSON(200,gin.H{
		"user_id": userID,
		"role":role,
	})
}

// Request
//  ↓
// AuthMiddleware
//    ├─ no token → BLOCK ❌
//    ├─ invalid token → BLOCK ❌
//    └─ valid token → c.Set(user_id, role)
//  ↓
// Profile controller
