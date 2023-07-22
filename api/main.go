package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// models

type User struct {
	Username string
}

type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Obtained bool   `json:"obtained"`
	Store    string `json:"store"`
	Aisle    int8   `json:"aisle"`
	Quantity int8   `json:"quantity"`
	Price    int64  `json:"price"`
	// Added_by User   `json:"added_by" gorm:"embedded"`
	//todo Picture
}

type Ingredient struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Obtained bool   `json:"obtained"`
	Store    string `json:"store"`
	Aisle    int8   `json:"aisle"`
	Quantity int8   `json:"quantity"`
	Price    int64  `json:"price"`
	Added_by User   `json:"added_by" gorm:"embedded"`
	//todo Picture
}

// collection of ingredients to add to a list
type Recipe struct {
	ID          int
	Ingredients Ingredient `json:"ingredients" gorm:"embedded"`
	Added_by    User       `gorm:"embedded"`
}

// collection to buy
type List struct {
	Items       Item       `gorm:"embedded"`
	Ingredients Ingredient `json:"ingredients" gorm:"embedded"`
	Added_by    User       `gorm:"embedded"`
}

//todo section out above to their own files

var testItems = []Item{
	{ID: 1, Name: "Test 1", Obtained: true, Store: "Krogers", Aisle: 4, Quantity: 5, Price: 5},
	{ID: 2, Name: "Test 2", Obtained: true, Store: "Krogers", Aisle: 4, Quantity: 5, Price: 5},
	{ID: 3, Name: "Test 3", Obtained: true, Store: "Krogers", Aisle: 4, Quantity: 5, Price: 5},
}

// CRUD
func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, testItems)
	// items := db.Find(&Items)
	// c.Data(http.StatusOK, items)
}

func createItem(c *gin.Context) {
	var newItem Item

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	testItems = append(testItems, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

func itemById(c *gin.Context) {
	id := c.Param("id")
	intid, err := strconv.Atoi(id)
	item, err := getItemById(intid)

	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, item)
}

func getItemById(id int) (*Item, error) {
	for i, I := range testItems {
		if I.ID == id {
			return &testItems[i], nil
		}
	}
	return nil, errors.New("Item not found")
}

func main() {

	//?env load
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	//?db connect
	var dbhost, dbuser, dbpass, dbname, dbport = os.Getenv("host"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("port")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbhost, dbuser, dbpass, dbname, dbport)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("unable to connect to database")
		panic("failed to connect to database")
	}
	db.AutoMigrate(&User{}, &Item{}, &Ingredient{}, &Recipe{}, &List{})

	//?Run app
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hurray!",
		})
	})
	router.GET("/item", getItems)
	router.GET("/item/:id", itemById)
	router.POST("/item", createItem)
	router.Run()
}
