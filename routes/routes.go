package routes

import (
	"time"
	"toap-go-backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // URL do frontend durante desenvolvimento
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/greeting/:name", controllers.Greeting)

	// Characters
	r.GET("/characters", controllers.AllCharacters)
	r.GET("/characters/:id", controllers.GetCharacter)
	r.POST("/characters", controllers.CreateCharacter)
	r.PUT("/characters/:id", controllers.UpdateCharacter)
	r.DELETE("/characters/:id", controllers.DeleteCharacter)
	r.POST("/characters/add-skills", controllers.AddSkillToCharacter)

	// Races
	r.GET("/races", controllers.AllRaces)
	r.GET("/races/:id", controllers.GetRace)
	r.POST("/races", controllers.CreateRace)
	r.PUT("/races/:id", controllers.UpdateRace)
	r.DELETE("/races/:id", controllers.DeleteRace)

	// Realms
	r.GET("/realms", controllers.AllRealms)
	r.GET("/realms/:id", controllers.GetRealm)
	r.POST("/realms", controllers.CreateRealm)
	r.PUT("/realms/:id", controllers.UpdateRealm)
	r.DELETE("/realms/:id", controllers.DeleteRealm)

	// Organizations
	r.GET("/organizations", controllers.AllOrganizations)
	r.GET("/organizations/:id", controllers.GetOrganization)
	r.POST("/organizations", controllers.CreateOrganization)
	r.PUT("/organizations/:id", controllers.UpdateOrganization)
	r.DELETE("/organizations/:id", controllers.DeleteOrganization)

	// Skills
	r.GET("/skills", controllers.AllSkills)
	r.GET("/skills/:id", controllers.GetSkill)
	r.POST("/skills", controllers.CreateSkill)
	r.PUT("/skills/:id", controllers.UpdateSkill)
	r.DELETE("/skills/:id", controllers.DeleteSkill)
	r.Run()
}
