package main

import (
	"toap-go-backend/database"
	"toap-go-backend/routes"
)

func main() {
	database.ConnectWithDB()
	database.MigrateDb()
	routes.HandleRequests()
}
