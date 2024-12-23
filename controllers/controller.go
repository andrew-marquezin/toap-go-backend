package controllers

import (
	"net/http"
	"toap-go-backend/database"
	"toap-go-backend/models"

	"github.com/gin-gonic/gin"
)

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"API says:": "Hi " + name + "! What's up?",
	})
}

func AllCharacters(c *gin.Context) {
	var characters []models.Character
	database.DB.Find(&characters)
	c.JSON(http.StatusOK, characters)
}
