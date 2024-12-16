package models

import "gorm.io/gorm"

type base struct {
	gorm.Model
	Name        string
	Description string
}

// type Image struct {}

type Realm struct {
	base
	Image string
}

type Organization struct {
	base
	Type string
}

type Skill struct {
	base
	Type string
}

type Race struct {
	base
	Story           string
	Characteristics string
}

type Character struct {
	base
	Image          string
	RaceID         int
	RealmID        int
	OrganizationID int
	Skills         []Skill
	// AppearsIn
	// RelatedCharacters
	// Category
}

type Scenario struct {
	base
	Image   string
	RealmID int
	Type    string
}

type Tale struct {
	Title      string
	Briefing   string
	Tale       string
	ScenarioID int
	Image      string
	// NotableCharacters
	// RelatedTales
}
