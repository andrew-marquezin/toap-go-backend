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
		"message": "Hi " + name + "! What's up?",
	})
}

// Character
func AllCharacters(c *gin.Context) {
	var characters []models.Character
	database.DB.Find(&characters)
	c.JSON(http.StatusOK, characters)
}

func GetCharacter(c *gin.Context) {
	var character models.Character
	id := c.Params.ByName("id")
	database.DB.First(&character, id)
	c.JSON(http.StatusOK, character)
}

func CreateCharacter(c *gin.Context) {
	var character models.Character

	if err := c.ShouldBindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&character)
	c.JSON(http.StatusOK, character)
}

// Race
func AllRaces(c *gin.Context) {
	var races []models.Race
	database.DB.Find(&races)
	c.JSON(http.StatusOK, races)
}

func GetRace(c *gin.Context) {
	var race models.Race
	id := c.Params.ByName("id")
	database.DB.First(&race, id)
	c.JSON(http.StatusOK, race)
}

func CreateRace(c *gin.Context) {
	var race models.Race

	if err := c.ShouldBindJSON(&race); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&race)
	c.JSON(http.StatusOK, race)
}

func UpdateRace(c *gin.Context) {
	var race models.Race
	id := c.Params.ByName("id")
	database.DB.First(&race, id)

	if err := c.ShouldBindJSON(&race); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&race).Updates(race)
}

func DeleteRace(c *gin.Context) {
	var race models.Race
	id := c.Params.ByName("id")

	database.DB.Delete(&race, id)
	c.JSON(http.StatusOK, gin.H{"message": "Race deleted successfully"})
}

// Realm
func AllRealms(c *gin.Context) {
	var realms []models.Realm
	database.DB.Find(&realms)
	c.JSON(http.StatusOK, realms)
}

func GetRealm(c *gin.Context) {
	var realm models.Realm
	id := c.Params.ByName("id")
	database.DB.First(&realm, id)
	c.JSON(http.StatusOK, realm)
}

func CreateRealm(c *gin.Context) {
	var realm models.Realm

	if err := c.ShouldBindJSON(&realm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&realm)
	c.JSON(http.StatusOK, realm)
}

func UpdateRealm(c *gin.Context) {
	var realm models.Realm
	id := c.Params.ByName("id")
	database.DB.First(&realm, id)

	if err := c.ShouldBindJSON(&realm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&realm).Updates(realm)
}

func DeleteRealm(c *gin.Context) {
	var realm models.Realm
	id := c.Params.ByName("id")

	database.DB.Delete(&realm, id)
	c.JSON(http.StatusOK, gin.H{"message": "Realm deleted successfully"})
}

// Organization
func AllOrganizations(c *gin.Context) {
	var organizations []models.Organization
	database.DB.Find(&organizations)
	c.JSON(http.StatusOK, organizations)
}

func GetOrganization(c *gin.Context) {
	var organization models.Organization
	id := c.Params.ByName("id")
	database.DB.First(&organization, id)
	c.JSON(http.StatusOK, organization)
}

func CreateOrganization(c *gin.Context) {
	var organization models.Organization

	if err := c.ShouldBindJSON(&organization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&organization)
	c.JSON(http.StatusOK, organization)
}

func UpdateOrganization(c *gin.Context) {
	var organization models.Organization
	id := c.Params.ByName("id")
	database.DB.First(&organization, id)

	if err := c.ShouldBindJSON(&organization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&organization).Updates(organization)
}

func DeleteOrganization(c *gin.Context) {
	var organization models.Organization
	id := c.Params.ByName("id")

	database.DB.Delete(&organization, id)
	c.JSON(http.StatusOK, gin.H{"message": "Organization deleted successfully"})
}
