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
	// for _, model := range []interface{}{&models.Character{}, &models.Tale{}, &models.Scenario{}, &models.Realm{}, &models.Organization{}} {
	// 	err = DB.AutoMigrate(model)
	// 	if err != nil {
	// 		log.Println("Erro ao migrar o modelo", model)
	// 		continue
	// 	}
	// }
	err = DB.AutoMigrate(&models.Character{}, &models.Skill{}, &models.Race{}, &models.Realm{}, &models.Organization{}, &models.Tale{}, &models.Scenario{})
	if err != nil {
		log.Panic("Erro na migração", err)
	}
}
