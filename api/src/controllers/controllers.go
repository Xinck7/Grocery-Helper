package controllers

import (
	"fmt"
	"net/http"

	"xinck/api/src/config"
	"xinck/api/src/models"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// Define database client
var db *gorm.DB = config.ConnectDB()

// ? requests section
type itemRequest struct {
	Name     string `json:"name"`
	Obtained bool   `json:"obtained"`
	Store    string `json:"store"`
	Aisle    int8   `json:"aisle"`
	Quantity int8   `json:"qty"`
	Price    int64  `json:"price"`
}

type ingredientRequest struct {
	Name     string `json:"name"`
	Obtained bool   `json:"obtained"`
	Store    string `json:"store"`
	Aisle    int8   `json:"aisle"`
	Quantity int8   `json:"qty"`
	Price    int64  `json:"price"`
}

type listRequest struct {
	Name        string              `json:"name"`
	Items       []models.Item       `json:"items"`
	Ingredients []models.Ingredient `json:"ingredients"`
	Recipes     []models.Recipe     `json:"recipes"`
}

type recipeRequest struct {
	Name        string              `json:"name"`
	Ingredients []models.Ingredient `json:"ingredients"`
}

// ? responses section
type itemResponse struct {
	itemRequest
	ID uint `json:"id"`
}

type ingredientResponse struct {
	ingredientRequest
	ID uint `json:"id"`
}

type listResponse struct {
	listRequest
	ID uint `json:"id"`
}

type recipeResponse struct {
	recipeRequest
	ID uint `json:"id"`
}

//! CRUD Methods

// ? Items section
func CreateItem(context *gin.Context) {
	var data itemRequest

	if err := context.Bind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := models.Item{}
	item.Name = data.Name
	item.Obtained = data.Obtained
	item.Store = data.Store
	item.Aisle = data.Aisle
	item.Quantity = data.Quantity
	item.Price = data.Price

	result := db.Create(&item)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong creating item to db"})
		return
	}

	var response itemResponse
	response.ID = item.ID
	response.Name = item.Name

	context.JSON(http.StatusCreated, response)
}

func GetAllItems(context *gin.Context) {
	var item []models.Item

	err := db.Find(&item)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting all items"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    item,
	})
}

func GetItemByID(context *gin.Context) {
	var data itemRequest

	reqParamId := context.Param("idItem")
	iditem := cast.ToUint(reqParamId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := models.Item{}

	itemById := db.Where("id = ?", iditem).First(&item)
	if itemById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "item not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    itemById,
	})
}

func UpdateItem(context *gin.Context) {
	var data itemRequest

	reqParamId := context.Param("idItem")
	iditem := cast.ToUint(reqParamId)

	item := models.Item{}

	itemById := db.Where("id = ?", iditem).First(&item)
	if itemById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "item not found"})
		return
	}

	item.Name = data.Name
	item.Obtained = data.Obtained
	item.Store = data.Store
	item.Aisle = data.Aisle
	item.Quantity = data.Quantity
	item.Price = data.Price

	result := db.Save(&item)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong updating item to db"})
		return
	}

	var response itemResponse
	response.ID = item.ID
	response.Name = item.Name
	response.Obtained = item.Obtained
	response.Store = item.Store
	response.Aisle = item.Aisle
	response.Quantity = item.Quantity
	response.Price = item.Price

	context.JSON(http.StatusCreated, response)
}

func DeleteItem(context *gin.Context) {
	item := models.Item{}

	reqParamId := context.Param("idItem")
	iditem := cast.ToUint(reqParamId)

	delete := db.Where("id = ?", iditem).Unscoped().Delete(&item)
	fmt.Println(delete)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    iditem,
	})
}

// ? Ingredients section
func CreateIngredient(context *gin.Context) {
	var data ingredientRequest

	if err := context.Bind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ingredient := models.Ingredient{}
	ingredient.Name = data.Name
	ingredient.Obtained = data.Obtained
	ingredient.Store = data.Store
	ingredient.Aisle = data.Aisle
	ingredient.Quantity = data.Quantity
	ingredient.Price = data.Price

	result := db.Create(&ingredient)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong creating ingredient to db"})
		return
	}

	var response ingredientResponse
	response.ID = ingredient.ID
	response.Name = ingredient.Name

	context.JSON(http.StatusCreated, response)
}

func GetAllIngredients(context *gin.Context) {
	var ingredient []models.Ingredient

	err := db.Find(&ingredient)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting ingredient from db"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    ingredient,
	})

}

func GetIngredientByID(context *gin.Context) {
	var data ingredientRequest

	reqParamId := context.Param("idIngredient")
	idingredient := cast.ToUint(reqParamId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ingredient := models.Ingredient{}

	ingredientById := db.Where("id = ?", idingredient).First(&ingredient)
	if ingredientById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ingredient not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    ingredientById,
	})
}

func UpdateIngredient(context *gin.Context) {
	var data ingredientRequest

	reqParamId := context.Param("idIngredient")
	idingredient := cast.ToUint(reqParamId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ingredient := models.Ingredient{}

	ingredientById := db.Where("id = ?", idingredient).First(&ingredient)
	if ingredientById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ingredient not found"})
		return
	}

	ingredient.Name = data.Name
	ingredient.Obtained = data.Obtained
	ingredient.Store = data.Store
	ingredient.Aisle = data.Aisle
	ingredient.Quantity = data.Quantity
	ingredient.Price = data.Price

	result := db.Save(&ingredient)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong updating ingredient to db"})
		return
	}

	var response ingredientResponse
	response.ID = ingredient.ID
	response.Name = ingredient.Name
	response.Obtained = ingredient.Obtained
	response.Store = ingredient.Store
	response.Aisle = ingredient.Aisle
	response.Quantity = ingredient.Quantity
	response.Price = ingredient.Price

	context.JSON(http.StatusCreated, response)
}

func DeleteIngredient(context *gin.Context) {
	ingredient := models.Ingredient{}

	reqParamId := context.Param("idIngredient")
	idingredient := cast.ToUint(reqParamId)

	delete := db.Where("id = ?", idingredient).Unscoped().Delete(&ingredient)
	fmt.Println(delete)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    idingredient,
	})
}

// ? Lists section
func CreateList(context *gin.Context) {
	var data listRequest

	if err := context.Bind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list := models.List{}
	list.Name = data.Name
	list.Items = data.Items
	list.Ingredients = data.Ingredients
	list.Recipes = data.Recipes

	result := db.Create(&list)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong creating list to db"})
		return
	}

	var response listResponse
	response.ID = list.ID
	response.Name = list.Name

	context.JSON(http.StatusCreated, response)
}

func GetAllLists(context *gin.Context) {
	var list []models.List

	err := db.Find(&list)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting list from db"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    list,
	})
}

func GetListByID(context *gin.Context) {
	var data listRequest

	reqParamId := context.Param("idList")
	idlist := cast.ToUint(reqParamId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list := models.List{}

	listById := db.Where("id = ?", idlist).First(&list)
	if listById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "list not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    listById,
	})
}

func UpdateList(context *gin.Context) {
	var data listRequest

	reqParamId := context.Param("idList")
	idlist := cast.ToUint(reqParamId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list := models.List{}

	listById := db.Where("id = ?", idlist).First(&list)
	if listById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "list not found"})
		return
	}

	list.Name = data.Name
	list.Items = data.Items
	list.Ingredients = data.Ingredients
	list.Recipes = data.Recipes

	result := db.Save(&list)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong updating list to db"})
		return
	}

	var response listResponse
	response.ID = list.ID
	response.Name = list.Name
	response.Items = list.Items
	response.Ingredients = list.Ingredients
	response.Recipes = list.Recipes

	context.JSON(http.StatusCreated, response)
}

func DeleteList(context *gin.Context) {
	list := models.List{}

	reqParamId := context.Param("idList")
	idlist := cast.ToUint(reqParamId)

	delete := db.Where("id = ?", idlist).Unscoped().Delete(&list)
	fmt.Println(delete)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    idlist,
	})
}

// ? Recipes section
func CreateRecipe(context *gin.Context) {
	var data recipeRequest

	if err := context.Bind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe := models.Recipe{}
	recipe.Name = data.Name
	recipe.Ingredients = data.Ingredients

	result := db.Create(&recipe)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong creating recipe to db"})
		return
	}

	var response recipeResponse
	response.ID = recipe.ID
	response.Name = recipe.Name

	context.JSON(http.StatusCreated, response)
}

func GetAllRecipes(context *gin.Context) {
	var recipe []models.Recipe

	err := db.Find(&recipe)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting recipe from db"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    recipe,
	})
}

func GetRecipeByID(context *gin.Context) {
	var data recipeRequest

	reqParamId := context.Param("idRecipe")
	idrecipe := cast.ToUint(reqParamId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe := models.Recipe{}

	recipeById := db.Where("id = ?", idrecipe).First(&recipe)
	if recipeById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "recipe not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    recipeById,
	})
}

func UpdateRecipe(context *gin.Context) {
	var data recipeRequest

	reqParamId := context.Param("idRecipe")
	idrecipe := cast.ToUint(reqParamId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe := models.Recipe{}

	recipeById := db.Where("id = ?", idrecipe).First(&recipe)
	if recipeById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "recipe not found"})
		return
	}

	recipe.Name = data.Name
	recipe.Ingredients = data.Ingredients

	result := db.Save(&recipe)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong updating recipe to db"})
		return
	}

	var response recipeResponse
	response.ID = recipe.ID
	response.Ingredients = recipe.Ingredients

	context.JSON(http.StatusCreated, response)
}

func DeleteRecipe(context *gin.Context) {

	recipe := models.Recipe{}

	reqParamId := context.Param("idRecipe")
	idrecipe := cast.ToUint(reqParamId)

	delete := db.Where("id = ?", idrecipe).Unscoped().Delete(&recipe)
	fmt.Println(delete)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    idrecipe,
	})
}
