package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"xinck/api/src/controllers"
	"xinck/api/src/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const itemsBaseRoute = "/items/"
const itemsIDRoute = "1"
const ingredientsBaseRoute = "/ingredients"
const ingredientsIDRoute = "/:id"
const listsBaseRoute = "/lists"
const listIDRoute = "/:id"
const recipesBaseRoute = "/recipes"
const recipesIDRoute = "/:id"

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// ? Items
func TestItemCreate(t *testing.T) {
	r := SetUpRouter()
	r.POST(itemsBaseRoute, controllers.CreateItem)

	item := models.Item{
		Name:     "test broom",
		Obtained: false,
		Aisle:    5,
		Store:    "Krogers",
		Quantity: 15,
		Price:    20,
	}
	jsonValue, _ := json.Marshal(item)
	req, _ := http.NewRequest("POST", itemsBaseRoute, bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

}

func TestItemGet(t *testing.T) {
	r := SetUpRouter()
	r.GET(itemsBaseRoute, controllers.GetAllItems)
	req, _ := http.NewRequest("GET", itemsBaseRoute, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

// func TestItemById(t *testing.T) {
// 	r := SetUpRouter()
// 	r.POST(itemsBaseRoute, controllers.CreateItem)

// 	item := models.Item{
// 		Name:     "test broom",
// 		Obtained: false,
// 		Aisle:    5,
// 		Store:    "Krogers",
// 		Quantity: 15,
// 		Price:    20,
// 	}
// 	jsonValue, _ := json.Marshal(item)
// 	req1, _ := http.NewRequest("POST", itemsBaseRoute, bytes.NewBuffer(jsonValue))
// 	w1 := httptest.NewRecorder()
// 	r.ServeHTTP(w1, req1)
// 	assert.Equal(t, 201, w1.Code)

// 	req, _ := http.NewRequest("GET", itemsBaseRoute+itemsIDRoute, nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, 200, w.Code)
// }

// func TestItemPut(t *testing.T) {

// }

// func TestItemDelete(t *testing.T) {

// }

//? Ingredients
// func TestIngredientCreate(t *testing.T) {

// }

func TestIngredientGet(t *testing.T) {
	r := SetUpRouter()
	r.GET(ingredientsBaseRoute, controllers.GetAllIngredients)
	req, _ := http.NewRequest("GET", ingredientsBaseRoute, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

// func TestIngredientById(t *testing.T) {

// }

// func TestIngredientPut(t *testing.T) {

// }

// func TestIngredientDelete(t *testing.T) {

// }

//? List
// func TestListCreate(t *testing.T) {

// }

func TestListGet(t *testing.T) {
	r := SetUpRouter()
	r.GET(listsBaseRoute, controllers.GetAllLists)
	req, _ := http.NewRequest("GET", listsBaseRoute, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

// func TestListById(t *testing.T) {

// }

// func TestListPut(t *testing.T) {

// }

// func TestListDelete(t *testing.T) {

// }

//? Recipe
// func TestRecipeCreate(t *testing.T) {

// }

func TestRecipeGet(t *testing.T) {
	r := SetUpRouter()
	r.GET(recipesBaseRoute, controllers.GetAllRecipes)
	req, _ := http.NewRequest("GET", recipesBaseRoute, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

// func TestRecipeById(t *testing.T) {

// }

// func TestRecipePut(t *testing.T) {

// }

// func TestRecipeDelete(t *testing.T) {

// }
