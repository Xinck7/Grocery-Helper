package models

import (
	"gorm.io/gorm"
)

// models
type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Username string
	Password string
}

type Item struct {
	gorm.Model
	ID       uint `gorm:"primary_key"`
	Name     string
	Obtained bool
	Store    string
	Aisle    int8
	Quantity int8
	Price    int64
	Lists    []List `gorm:"many2many:item_lists;"`
	Added    string
	// todo Picture
}

type Ingredient struct {
	gorm.Model
	ID       uint `gorm:"primary_key"`
	Name     string
	Obtained bool
	Store    string
	Aisle    int8
	Quantity int8
	Price    int64
	Lists    []List   `gorm:"many2many:ingredient_lists;"`
	Recipes  []Recipe `gorm:"many2many:ingredient_recipes;"`
	Added    string
	//todo Picture
}

type List struct {
	gorm.Model
	ID          uint `gorm:"primary_key"`
	Name        string
	Items       []Item       `gorm:"many2many:item_lists;"`
	Ingredients []Ingredient `gorm:"many2many:ingredient_lists;"`
	Recipes     []Recipe     `gorm:"many2many:list_recipes;"`
	Price       int64
	Added       string
}

type Recipe struct {
	gorm.Model
	ID          uint `gorm:"primary_key"`
	Name        string
	Ingredients []Ingredient `gorm:"many2many:ingredient_recipes;"`
	Lists       []List       `gorm:"many2many:list_recipes;"`
	Price       int64
	Added       string
}
