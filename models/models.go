package models

import "gorm.io/gorm"

// // db.Model(&models.UserInfo{}).AddForeignKey("u_id", "t_user(id)", "RESTRICT", "RESTRICT")

type Character struct {
	gorm.Model
	Name           string
	Description    string
	RaceID         uint
	Race           Race
	RealmID        uint
	Realm          Realm
	OrganizationID uint
	Organization   Organization
	Skills         []Skill `gorm:"many2many:character_skills;"`
	Tales          []Tale  `gorm:"many2many:character_tales;"`
	// Image       string
	// RelatedCharacters
	// Category
}

type Tale struct {
	gorm.Model
	Title      string
	Briefing   string
	Content    string
	ScenarioID uint
	Scenario   Scenario
	Image      string
	Author     string
	// NotableCharacters []Character `gorm:"many2many:character_tale"`
	// RelatedTales
}

type Scenario struct {
	gorm.Model
	Image   string
	RealmID uint
	Realm   Realm
	Type    string
}

type Realm struct {
	gorm.Model
	Name        string
	Description string
	Image       string
	// Scenarios []Scenario `gorm:"foreignKey:RealmID"` // 1:N
}

type Organization struct {
	gorm.Model
	Name        string
	Description string
	// Type string
	// Characters []Character `gorm:"foreignKey:OrganizationID"` // 1:N
}

type Skill struct {
	gorm.Model
	Name        string
	Description string
	// Characters []Character `gorm:"many2many:character_skill"`
}

type Race struct {
	gorm.Model
	Name            string
	Story           string
	Characteristics string
	// Characters      []Character `gorm:"foreignKey:RaceID"` // 1:N
}
