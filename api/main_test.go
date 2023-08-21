package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"xinck/api/controllers"
	"xinck/api/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const itemsBaseRoute = "/items"

// const itemsIDRoute = "/:id"

const ingredientsBaseRoute = "/ingredients"

// const ingredientsIDRoute = "/:id"

const listsBaseRoute = "/lists"

// const listIDRoute = "/:id"

const recipesBaseRoute = "/recipes"

// const recipesIDRoute = "/:id"

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) //comment for test debug
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
// 	r.GET(itemsBaseRoute+itemsIDRoute, controllers.GetItemByID)

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

// ? Ingredients
func TestIngredientCreate(t *testing.T) {
	r := SetUpRouter()
	r.POST(ingredientsBaseRoute, controllers.CreateIngredient)

	ingredient := models.Ingredient{
		Name:     "test eggs",
		Obtained: true,
		Aisle:    5,
		Store:    "Krogers",
		Quantity: 15,
		Price:    20,
	}
	jsonValue, _ := json.Marshal(ingredient)
	req, _ := http.NewRequest("POST", ingredientsBaseRoute, bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

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

// ? List
func TestListCreate(t *testing.T) {
	r := SetUpRouter()
	r.POST(listsBaseRoute, controllers.CreateList)
	item := models.Item{
		Name:     "test broom",
		Obtained: false,
		Aisle:    5,
		Store:    "Krogers",
		Quantity: 15,
		Price:    20,
	}
	item1 := models.Item{
		Name:     "test broom band",
		Obtained: false,
		Aisle:    5,
		Store:    "Krogers",
		Quantity: 15,
		Price:    20,
	}
	ingredient := models.Ingredient{
		Name:     "test eggs",
		Obtained: true,
		Aisle:    5,
		Store:    "Krogers",
		Quantity: 15,
		Price:    20,
	}
	ingredient1 := models.Ingredient{
		Name:     "test bacon",
		Obtained: true,
		Aisle:    5,
		Store:    "Krogers",
		Quantity: 15,
		Price:    20,
	}
	list := models.List{
		Name:        "test eggs",
		Items:       []models.Item{item, item1},
		Ingredients: []models.Ingredient{ingredient, ingredient1},
		Price:       20,
	}
	jsonValue, _ := json.Marshal(list)
	req, _ := http.NewRequest("POST", listsBaseRoute, bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

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

// ? Recipe
func TestRecipeCreate(t *testing.T) {
	r := SetUpRouter()
	r.POST(recipesBaseRoute, controllers.CreateRecipe)
	ingredient := models.Ingredient{
		Name:     "test eggs",
		Obtained: true,
		Aisle:    5,
		Store:    "Krogers",
		Quantity: 15,
		Price:    20,
	}
	ingredient1 := models.Ingredient{
		Name:     "test bacon",
		Obtained: true,
		Aisle:    5,
		Store:    "Krogers",
		Quantity: 15,
		Price:    20,
	}
	list := models.List{
		Name:        "test eggs",
		Ingredients: []models.Ingredient{ingredient, ingredient1},
		Price:       20,
	}
	jsonValue, _ := json.Marshal(list)
	req, _ := http.NewRequest("POST", recipesBaseRoute, bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

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
