package routes

import (
	"toap-go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/greeting/:name", controllers.Greeting)
	r.GET("/characters", controllers.AllCharacters)
	r.Run()
}
