package models

import (
	"gorm.io/gorm"
)

// models
type User struct {
	gorm.Model
	Username string
	// Items []  //eg - `gorm:"foreignKey:UserID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Item struct {
	gorm.Model
	Name     string
	Obtained bool
	Store    string
	Aisle    int8
	Quantity int8
	Price    int64
	// Added_by User.Username
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
	// Added_by User
	//todo Picture
}

// collection to buy
type List struct {
	gorm.Model
	Name        string
	Items       []Item       `gorm:"many2many:list_items;"`
	Ingredients []Ingredient `gorm:"many2many:list_ingredients;"`
	Recipes     []Recipe     `gorm:"many2many:list_recipes;"`
	Price       int64
	// Added_by    User
}

// collection of ingredients to add to a list
type Recipe struct {
	gorm.Model
	Name        string
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients;"`
	Price       int64
	// Added_by    User
}
