package models

import (
	"gorm.io/gorm"
)

// models
type User struct {
	gorm.Model
	Username string
}

type Item struct {
	gorm.Model
	Name     string
	Obtained bool
	Store    string
	Aisle    int8
	Quantity int8
	Price    int64
	// Added_by User
	// todo Picture
}

type Ingredient struct {
	gorm.Model
	Name     string
	Obtained bool
	Store    string
	Aisle    int8
	Quantity int8
	Price    int64
	// Added_by User `gorm:"embedded"`
	//todo Picture
}

// collection of ingredients to add to a list
type Recipe struct {
	gorm.Model
	Ingredients Ingredient `gorm:"embedded"`
	// Added_by    User       `gorm:"embedded"`
}

// collection to buy
type List struct {
	Items       Item       `gorm:"embedded"`
	Ingredients Ingredient `gorm:"embedded"`
	// Added_by    User       `gorm:"embedded"`
}
