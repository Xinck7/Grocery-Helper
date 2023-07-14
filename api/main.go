package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

// models 

type User struct(
	username	string `json:"username"`
)

type Item struct(
	gorm.Model
	ID			string `json:"id"`
	Name		string `json:"name"`
	Obtained	bool `json:"obtained"`
	Store		string `json:"store"`
	Aisle		string `json:"aisle"`
	Quantity	int8 `json:"quantity"`
	Price		int64 `json:"price"`
	Added_by	User `json:"added_by"`
	//todo Picture
)

type Ingredient struct(
	gorm.Model
	ID			string `json:"id"`
	Name		string `json:"name"`
	Obtained	bool `json:"obtained"`
	Store		string `json:"store"`
	Aisle		string `json:"aisle"`
	Quantity	int8 `json:"quantity"`
	Price		int64 `json:"price"`
	Added_by	User `json:"added_by"`
	//todo Picture
)

// collection of ingredients to add to a list
type Recipe struct(
	gorm.Model
	Ingredients Ingredient
	Added_by User
)

// collection to buy
type List struct(
	gorm.Model
	Item Items
	Ingredients Ingredient
)



//todo define postgres connection
//todo section out above to their own files

func main(){
	router := gin.Default()
	router.Run("localhost:8080")
}