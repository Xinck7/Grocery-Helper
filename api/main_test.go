package main_test

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
const ingredientsBaseRoute = "/ingredients"
const listsBaseRoute = "/lists"
const recipesBaseRoute = "/recipes"

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
	recipe := models.Recipe{
		Name:        "test recipe",
		Ingredients: []models.Ingredient{ingredient, ingredient1},
		Price:       20,
	}
	list := models.List{
		Name:        "test list",
		Items:       []models.Item{item, item1},
		Ingredients: []models.Ingredient{ingredient, ingredient1},
		Recipes:     []models.Recipe{recipe},
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
	recipe := models.Recipe{
		Name:        "test recipe",
		Ingredients: []models.Ingredient{ingredient, ingredient1},
		Price:       20,
	}
	jsonValue, _ := json.Marshal(recipe)
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
