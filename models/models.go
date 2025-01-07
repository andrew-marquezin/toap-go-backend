package models

import (
	"time"

	"gorm.io/gorm"
)

// // db.Model(&models.UserInfo{}).AddForeignKey("u_id", "t_user(id)", "RESTRICT", "RESTRICT")

type Character struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	Name           string         `json:"name"`
	RaceID         uint           `json:"raceID"`
	Race           Race           `gorm:"foreignKey:RaceID" json:"race"`
	Skills         []Skill        `gorm:"many2many:character_skills;" json:"skills"` // Many-to-Many com Skill
	OrganizationID uint           `json:"organizationID"`
	Organization   Organization   `gorm:"foreignKey:OrganizationID" json:"organization"`
	Tales          []Tale         `gorm:"many2many:character_tales;" json:"tales"` // Many-to-Many com Tale
	RealmID        uint           `json:"realmID"`
	Realm          Realm          `gorm:"foreignKey:RealmID" json:"realm"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	// Description    string
	// Image       string
	// Category
}

type Tale struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Title      string         `json:"title"`
	Synopsis   string         `json:"synopsis"`
	Content    string         `json:"content"`
	ScenarioID uint           `json:"scenarioID"`                                   // Chave estrangeira para Scenario
	Scenario   Scenario       `gorm:"foreignKey:ScenarioID" json:"scenario"`        // Referência ao Scenario
	Characters []Character    `gorm:"many2many:character_tales;" json:"characters"` // Many-to-Many com Character
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Scenario struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Tales       []Tale         `gorm:"foreignKey:ScenarioID" json:"tales"` // One-to-Many com Tale
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	// Type      string
}

type Realm struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Characters  []Character    `gorm:"foreignKey:RealmID" json:"characters"` // One-to-Many com Character
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	// Scenarios []Scenario `gorm:"foreignKey:RealmID"` // 1:N
}

type Organization struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Purpose   string         `json:"purpose"`
	Members   []Character    `gorm:"foreignKey:OrganizationID" json:"members"` // One-to-Many com Character
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Skill struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Characters  []Character    `gorm:"many2many:character_skills;" json:"characters"` // Many-to-Many com Character
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Race struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Name       string         `json:"name"`
	Story      string         `json:"story"`
	Traits     string         `json:"traits"`
	Characters []Character    `gorm:"foreignKey:RaceID" json:"raceID"` // 1:N
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
