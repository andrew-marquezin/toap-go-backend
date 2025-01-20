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

func contains(ids []uint, id uint) bool {
	for _, v := range ids {
		if v == id {
			return true
		}
	}
	return false
}

// Character
func AllCharacters(c *gin.Context) {
	var characters []models.Character
	if err := database.DB.
		Preload("Race").
		Preload("Organization").
		Preload("Realm").
		Preload("Skills").
		Find(&characters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, characters)
}

func GetCharacter(c *gin.Context) {
	var character models.Character
	id := c.Params.ByName("id")
	if err := database.DB.
		Preload("Race").
		Preload("Organization").
		Preload("Realm").
		Preload("Skills").
		First(&character, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}
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

func UpdateCharacter(c *gin.Context) {
	var character models.Character
	id := c.Params.ByName("id")

	if err := database.DB.Preload("Skills").First(&character, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}

	// Lê os dados enviados no body
	var input models.Character
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Atualiza os campos do personagem
	character.Name = input.Name
	character.RaceID = input.RaceID
	character.OrganizationID = input.OrganizationID
	character.RealmID = input.RealmID

	// Obter IDs das skills enviadas no payload
	var inputSkillIDs []uint
	for _, skill := range input.Skills {
		inputSkillIDs = append(inputSkillIDs, skill.ID)
	}

	// Identificar quais skills remover
	var skillsToRemove []models.Skill
	for _, skill := range character.Skills {
		if !contains(inputSkillIDs, skill.ID) {
			skillsToRemove = append(skillsToRemove, skill)
		}
	}

	// Atualizar skills: remover as antigas e adicionar as novas
	var skillsToAdd []models.Skill
	if len(input.Skills) > 0 {
		for _, skill := range input.Skills {
			var tempSkill models.Skill
			if err := database.DB.First(&tempSkill, skill.ID).Error; err == nil {
				skillsToAdd = append(skillsToAdd, tempSkill)
			}
		}
	}

	// Remove skills obsoletas
	if len(skillsToRemove) > 0 {
		database.DB.Model(&character).Association("Skills").Delete(skillsToRemove)
	}

	// Adiciona skills novas
	if len(skillsToAdd) > 0 {
		database.DB.Model(&character).Association("Skills").Replace(skillsToAdd)
	}

	// Salva as mudanças
	if err := database.DB.Save(&character).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, character)
}

func DeleteCharacter(c *gin.Context) {
	var character models.Character
	id := c.Params.ByName("id")

	database.DB.Delete(&character, id)
	c.JSON(http.StatusOK, gin.H{"message": "Character deleted successfully"})
}

func AddSkillToCharacter(c *gin.Context) {
	var request struct {
		CharacterID uint   `json:"character_id"`
		SkillIDs    []uint `json:"skill_ids"`
	}

	// Faz o bind do JSON recebido
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Busca o Character pelo ID
	var character models.Character
	if err := database.DB.First(&character, request.CharacterID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}

	// Busca as Skills pelos IDs
	var skills []models.Skill
	if err := database.DB.Find(&skills, request.SkillIDs).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "One or more skills not found"})
		return
	}

	// Adiciona as Skills ao Character
	if err := database.DB.Model(&character).Association("Skills").Append(&skills); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add skills"})
		return
	}

	// Retorna o Character atualizado
	c.JSON(http.StatusOK, gin.H{
		"message":   "Skills added successfully",
		"character": character,
	})
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

// Skill
func AllSkills(c *gin.Context) {
	var skills []models.Skill
	database.DB.Find(&skills)
	c.JSON(http.StatusOK, skills)
}

func GetSkill(c *gin.Context) {
	var skill models.Skill
	id := c.Params.ByName("id")
	database.DB.First(&skill, id)
	c.JSON(http.StatusOK, skill)
}

func CreateSkill(c *gin.Context) {
	var skill models.Skill

	if err := c.ShouldBindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&skill)
	c.JSON(http.StatusOK, skill)
}

func UpdateSkill(c *gin.Context) {
	var skill models.Skill
	id := c.Params.ByName("id")
	database.DB.First(&skill, id)

	if err := c.ShouldBindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&skill).Updates(skill)
	c.JSON(http.StatusOK, skill)
}

func DeleteSkill(c *gin.Context) {
	var skill models.Skill
	id := c.Params.ByName("id")

	database.DB.Delete(&skill, id)
	c.JSON(http.StatusOK, gin.H{"message": "Skill deleted successfully"})
}
