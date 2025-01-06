package models

import (
	"time"

	"gorm.io/gorm"
)

// // db.Model(&models.UserInfo{}).AddForeignKey("u_id", "t_user(id)", "RESTRICT", "RESTRICT")

type Character struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	RaceID         uint
	Race           Race    `gorm:"foreignKey:RaceID"`
	Skills         []Skill `gorm:"many2many:character_skills;"` // Many-to-Many com Skill
	OrganizationID uint
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	Tales          []Tale       `gorm:"many2many:character_tales;"` // Many-to-Many com Tale
	RealmID        uint
	Realm          Realm `gorm:"foreignKey:RealmID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	// Description    string
	// Image       string
	// Category
}

type Tale struct {
	ID         uint `gorm:"primaryKey"`
	Title      string
	Synopsis   string
	Content    string
	ScenarioID uint        // Chave estrangeira para Scenario
	Scenario   Scenario    `gorm:"foreignKey:ScenarioID"`      // Referência ao Scenario
	Characters []Character `gorm:"many2many:character_tales;"` // Many-to-Many com Character
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type Scenario struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Tales       []Tale `gorm:"foreignKey:ScenarioID"` // One-to-Many com Tale
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	// Type      string
}

type Realm struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Characters  []Character `gorm:"foreignKey:RealmID"` // One-to-Many com Character
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	// Scenarios []Scenario `gorm:"foreignKey:RealmID"` // 1:N
}

type Organization struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Purpose   string
	Members   []Character `gorm:"foreignKey:OrganizationID"` // One-to-Many com Character
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Skill struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Characters  []Character `gorm:"many2many:character_skills;"` // Many-to-Many com Character
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Race struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Story      string
	Traits     string
	Characters []Character `gorm:"foreignKey:RaceID"` // 1:N
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
