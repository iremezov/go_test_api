package controllers

import (
	"awesomeProject/helpers"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserInput struct {
	Login string `json:"login" binding:"required"`
	Email string `json:"email" binding:"required"`
	State int    `json:"state"`
}

// GET /books
// Get all books
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Login: input.Login, Email: input.Email, State: input.State}
	models.DB.Create(&user)

	u, _ := json.Marshal(user)
	helpers.Log_msg("CreateUser", string(u))

	c.JSON(http.StatusOK, gin.H{"data": user})

}

func FindUser(c *gin.Context) { // Get model if exist
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	u, _ := json.Marshal(user)
	helpers.Log_msg("UpdateUser", string(u))

	c.JSON(http.StatusOK, gin.H{"data": user})
}
