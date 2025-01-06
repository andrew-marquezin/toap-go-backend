package database

import (
	"log"
	"toap-go-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectWithDB() {
	connStr := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connStr))
	if err != nil {
		log.Panic("erro de conexão", err)
	}
}

func MigrateDb() {
	err = DB.AutoMigrate(&models.Character{}, &models.Skill{}, &models.Race{}, &models.Realm{}, &models.Organization{}, &models.Tale{}, &models.Scenario{})
	if err != nil {
		log.Panic("Erro na migração", err)
	}
}
