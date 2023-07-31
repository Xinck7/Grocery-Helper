package controllers

import (
	"net/http"

	"xinck/api/src/config"
	"xinck/api/src/models"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// Define database client
var db *gorm.DB = config.ConnectDB()

// item struct for request body
type itemRequest struct {
	Name     string `json:"name"`
	Obtained bool   `json:"obtained"`
	Store    string `json:"store"`
	Aisle    int8   `json:"aisle"`
	Quantity int8   `json:"qty"`
	Price    int64  `json:"price"`
}

// Defining struct for response
type itemResponse struct {
	itemRequest
	ID uint `json:"id"`
}

func CreateItem(context *gin.Context) {
	var data itemRequest

	// Binding request body json to request body struct
	if err := context.Bind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Matching item models struct with item request struct
	item := models.Item{}
	item.Name = data.Name
	item.Obtained = data.Obtained
	item.Store = data.Store
	item.Aisle = data.Aisle
	item.Quantity = data.Quantity
	item.Price = data.Price

	// Querying to database
	result := db.Create(&item)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong creating item to db"})
		return
	}

	// Matching result to create response
	var response itemResponse
	response.ID = item.ID
	response.Name = item.Name

	// Creating http response
	context.JSON(http.StatusCreated, response)
}

// Getting all item data
func GetAllItems(context *gin.Context) {
	var item []models.Item

	// Querying to find item data.
	err := db.Find(&item)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting all items"})
		return
	}

	// Creating http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    item,
	})

}

// Update item data
func UpdateItem(context *gin.Context) {
	var data itemRequest

	// Defining request parameter to get item id
	reqParamId := context.Param("iditem")
	iditem := cast.ToUint(reqParamId)

	// Binding request body json to request body struct
	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Initiate models item
	item := models.Item{}

	// Querying find item data by item id from request parameter
	itemById := db.Where("id = ?", iditem).First(&item)
	if itemById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "item not found"})
		return
	}

	// Matching item request with item models
	item.Name = data.Name
	item.Obtained = data.Obtained
	item.Store = data.Store
	item.Aisle = data.Aisle
	item.Quantity = data.Quantity
	item.Price = data.Price

	// Update new item data
	result := db.Save(&item)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong updating item to db"})
		return
	}

	// Matching result to item response struct
	var response itemResponse
	response.ID = item.ID
	response.Name = item.Name
	response.Obtained = item.Obtained
	response.Store = item.Store
	response.Aisle = item.Aisle
	response.Quantity = item.Quantity
	response.Price = item.Price

	// Creating http response
	context.JSON(http.StatusCreated, response)
}

// ingredient struct for request body
type ingredientRequest struct {
	Name     string `json:"name"`
	Obtained bool   `json:"obtained"`
	Store    string `json:"store"`
	Aisle    int8   `json:"aisle"`
	Quantity int8   `json:"qty"`
	Price    int64  `json:"price"`
}

// Defining struct for response
type ingredientResponse struct {
	ingredientRequest
	ID uint `json:"id"`
}

func CreateIngredient(context *gin.Context) {
	var data ingredientRequest

	// Binding request body json to request body struct
	if err := context.Bind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Matching ingredient models struct with ingredient request struct
	ingredient := models.Ingredient{}
	ingredient.Name = data.Name
	ingredient.Obtained = data.Obtained
	ingredient.Store = data.Store
	ingredient.Aisle = data.Aisle
	ingredient.Quantity = data.Quantity
	ingredient.Price = data.Price

	// Querying to database
	result := db.Create(&ingredient)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong creating ingredient to db"})
		return
	}

	// Matching result to create response
	var response ingredientResponse
	response.ID = ingredient.ID
	response.Name = ingredient.Name

	// Creating http response
	context.JSON(http.StatusCreated, response)
}

// Getting all ingredient data
func GetAllIngredients(context *gin.Context) {
	var ingredient []models.Ingredient

	// Querying to find ingredient data.
	err := db.Find(&ingredient)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting ingredient from db"})
		return
	}

	// Creating http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    ingredient,
	})

}

func UpdateIngredient(context *gin.Context) {
	var data ingredientRequest

	// Defining request parameter to get ingredient id
	reqParamId := context.Param("idingredient")
	idingredient := cast.ToUint(reqParamId)

	// Binding request body json to request body struct
	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Initiate models ingredient
	ingredient := models.Ingredient{}

	// Querying find ingredient data by ingredient id from request parameter
	ingredientById := db.Where("id = ?", idingredient).First(&ingredient)
	if ingredientById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ingredient not found"})
		return
	}

	// Matching ingredient request with ingredient models
	ingredient.Name = data.Name
	ingredient.Obtained = data.Obtained
	ingredient.Store = data.Store
	ingredient.Aisle = data.Aisle
	ingredient.Quantity = data.Quantity
	ingredient.Price = data.Price

	// Update new ingredient data
	result := db.Save(&ingredient)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong updating ingredient to db"})
		return
	}

	// Matching result to ingredient response struct
	var response ingredientResponse
	response.ID = ingredient.ID
	response.Name = ingredient.Name
	response.Obtained = ingredient.Obtained
	response.Store = ingredient.Store
	response.Aisle = ingredient.Aisle
	response.Quantity = ingredient.Quantity
	response.Price = ingredient.Price

	// Creating http response
	context.JSON(http.StatusCreated, response)
}
