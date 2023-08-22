package controllers

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"time"

	"xinck/api/config"
	"xinck/api/models"
	"xinck/api/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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
	Price       int64               `json:"price"`
}

type recipeRequest struct {
	Name        string              `json:"name"`
	Ingredients []models.Ingredient `json:"ingredients"`
	Price       int64               `json:"price"`
}

type registerRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string
}

type userRequest struct {
	Email    string `json:"email"`
	Password string
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

type registerResponse struct {
	registerRequest
	ID uint `json:"id"`
}

type userResponse struct {
	userRequest
	ID uint `json:"id"`
}

//! CRUD Methods

// ? Items section
func CreateItem(c *gin.Context) {
	var data itemRequest

	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong creating item to db"})
		return
	}

	var response itemResponse
	response.ID = item.ID
	response.Name = item.Name
	response.Store = item.Store
	response.Aisle = item.Aisle
	response.Quantity = item.Quantity
	response.Price = item.Price

	c.JSON(http.StatusCreated, response)
}

func GetAllItems(c *gin.Context) {
	var item []models.Item

	// todo make this all by user not full db
	err := db.Find(&item)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting all items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    item,
	})
}

func GetItemByID(c *gin.Context) {

	reqParamId := c.Param("id")
	iditem := cast.ToUint(reqParamId)
	item := models.Item{}

	itemById := db.Where("id = ?", iditem).First(&item)
	if itemById.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "item not found"})
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

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    response,
	})
}

func UpdateItem(c *gin.Context) {
	var data itemRequest

	reqParamId := c.Param("id")
	iditem := cast.ToUint(reqParamId)

	item := models.Item{}

	itemById := db.Where("id = ?", iditem).First(&item)
	if itemById.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "item not found"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong updating item to db"})
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

	c.JSON(http.StatusCreated, response)
}

func DeleteItem(c *gin.Context) {
	item := models.Item{}

	reqParamId := c.Param("id")
	iditem := cast.ToUint(reqParamId)

	delete := db.Where("id = ?", iditem).Unscoped().Delete(&item)
	fmt.Println(delete)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    iditem,
	})
}

// ? Ingredients section
func CreateIngredient(c *gin.Context) {
	var data ingredientRequest

	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong creating ingredient to db"})
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

	c.JSON(http.StatusCreated, response)
}

func GetAllIngredients(c *gin.Context) {
	var ingredient []models.Ingredient

	// todo make this all by user not full db
	err := db.Find(&ingredient)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting ingredient from db"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    ingredient,
	})

}

func GetIngredientByID(c *gin.Context) {
	reqParamId := c.Param("id")
	idIngredient := cast.ToUint(reqParamId)
	ingredient := models.Ingredient{}

	ingredientById := db.Where("id = ?", idIngredient).First(&ingredient)
	if ingredientById.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ingredient not found"})
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

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    response,
	})
}

func UpdateIngredient(c *gin.Context) {
	var data ingredientRequest

	reqParamId := c.Param("id")
	idIngredient := cast.ToUint(reqParamId)

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ingredient := models.Ingredient{}

	ingredientById := db.Where("id = ?", idIngredient).First(&ingredient)
	if ingredientById.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ingredient not found"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong updating ingredient to db"})
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

	c.JSON(http.StatusCreated, response)
}

func DeleteIngredient(c *gin.Context) {
	ingredient := models.Ingredient{}

	reqParamId := c.Param("id")
	idIngredient := cast.ToUint(reqParamId)

	delete := db.Where("id = ?", idIngredient).Unscoped().Delete(&ingredient)
	fmt.Println(delete)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    idIngredient,
	})
}

// ? Lists section
func CreateList(c *gin.Context) {
	var data listRequest

	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list := models.List{}
	list.Name = data.Name
	list.Items = data.Items
	list.Ingredients = data.Ingredients
	list.Recipes = data.Recipes

	result := db.Create(&list)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong creating list to db"})
		return
	}

	// todo foreach loop to get and append db info from the getbyID functions

	var response listResponse
	response.ID = list.ID
	response.Name = list.Name
	response.Items = list.Items
	response.Ingredients = list.Ingredients
	response.Recipes = list.Recipes

	c.JSON(http.StatusCreated, response)
}

func GetAllLists(c *gin.Context) {
	var list []models.List

	// todo make this all by user not full db
	err := db.Find(&list)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting list from db"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    list,
	})
}

func GetListByID(c *gin.Context) {
	reqParamId := c.Param("id")
	idList := cast.ToUint(reqParamId)

	list := models.List{}

	listById := db.Preload("Items").Preload("Ingredients").Preload("Recipes.Ingredients").First(&list, idList)
	if listById.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "list not found"})
		return
	}

	//calculates and updates the entry when checked
	listPrice := utils.CalculateListPrice(&list)
	list.Price = listPrice
	result := db.Save(&list)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong updating list to db"})
		return
	}

	var response listResponse
	response.ID = list.ID
	response.Name = list.Name
	response.Items = list.Items
	response.Ingredients = list.Ingredients
	response.Recipes = list.Recipes
	response.Price = list.Price

	sort.Slice(list.Items, func(i, j int) bool {
		return list.Items[i].Aisle < list.Items[j].Aisle
	})
	sort.Slice(list.Ingredients, func(i, j int) bool {
		return list.Ingredients[i].Aisle < list.Ingredients[j].Aisle
	})

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    response,
	})
}

func UpdateList(c *gin.Context) {
	var data listRequest

	reqParamId := c.Param("id")
	idList := cast.ToUint(reqParamId)

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list := models.List{}

	listById := db.Where("id = ?", idList).First(&list)
	if listById.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "list not found"})
		return
	}

	list.Name = data.Name
	list.Items = data.Items
	list.Ingredients = data.Ingredients
	list.Recipes = data.Recipes

	result := db.Save(&list)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong updating list to db"})
		return
	}

	var response listResponse
	response.ID = list.ID
	response.Name = list.Name
	response.Items = list.Items
	response.Ingredients = list.Ingredients
	response.Recipes = list.Recipes

	c.JSON(http.StatusCreated, response)
}

func DeleteList(c *gin.Context) {
	list := models.List{}

	reqParamId := c.Param("id")
	idList := cast.ToUint(reqParamId)

	delete := db.Where("id = ?", idList).Unscoped().Delete(&list)
	fmt.Println(delete)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    idList,
	})
}

// ? Recipes section
func CreateRecipe(c *gin.Context) {
	var data recipeRequest

	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe := models.Recipe{}
	recipe.Name = data.Name
	recipe.Ingredients = data.Ingredients

	result := db.Create(&recipe)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong creating recipe to db"})
		return
	}

	// todo foreach loop to get and append db info from the getbyID functions

	var response recipeResponse
	response.ID = recipe.ID
	response.Name = recipe.Name
	response.Ingredients = recipe.Ingredients

	c.JSON(http.StatusCreated, response)
}

func GetAllRecipes(c *gin.Context) {
	var recipe []models.Recipe

	// todo make this all by user not full db
	err := db.Find(&recipe)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting recipe from db"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    recipe,
	})
}

func GetRecipeByID(c *gin.Context) {
	reqParamId := c.Param("id")
	idRecipe := cast.ToUint(reqParamId)
	recipe := models.Recipe{}

	recipeById := db.Preload("Ingredients").First(&recipe, idRecipe)
	if recipeById.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "recipe not found"})
		return
	}

	//calculates and updates the entry when checked
	recipePrice := utils.CalculateRecipePrice(&recipe)
	recipe.Price = recipePrice
	result := db.Save(&recipe)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong updating list to db"})
		return
	}

	var response recipeResponse
	response.ID = recipe.ID
	response.Ingredients = recipe.Ingredients
	response.Price = recipe.Price

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    response,
	})
}

func UpdateRecipe(c *gin.Context) {
	var data recipeRequest

	reqParamId := c.Param("id")
	idRecipe := cast.ToUint(reqParamId)

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe := models.Recipe{}

	recipeById := db.Where("id = ?", idRecipe).First(&recipe)
	if recipeById.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "recipe not found"})
		return
	}

	recipe.Name = data.Name
	recipe.Ingredients = data.Ingredients

	result := db.Save(&recipe)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong updating recipe to db"})
		return
	}

	var response recipeResponse
	response.ID = recipe.ID
	response.Name = recipe.Name
	response.Ingredients = recipe.Ingredients

	c.JSON(http.StatusCreated, response)
}

func DeleteRecipe(c *gin.Context) {

	recipe := models.Recipe{}

	reqParamId := c.Param("id")
	idRecipe := cast.ToUint(reqParamId)

	delete := db.Where("id = ?", idRecipe).Unscoped().Delete(&recipe)
	fmt.Println(delete)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    idRecipe,
	})
}

// ? Users section
func RegisterUser(c *gin.Context) {
	var data registerRequest

	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong encrypting the password"})
		return
	}
	user := models.User{}
	user.Email = data.Email
	user.Username = data.Username
	user.Password = string(hashedPassword)

	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong creating user in db"})
		return
	}

	var response registerResponse
	response.ID = user.ID
	response.Email = user.Email
	response.Username = user.Username

	c.JSON(http.StatusCreated, response)
}

func LoginUser(c *gin.Context) {
	var data userRequest
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	userFromDb := db.First(&user, "email = ?", data.Email)
	if userFromDb.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ingredient not found"})
		return
	}

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 1).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN_SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func ValidateUserSession(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func LogoutUser(c *gin.Context) {
	user, _ := c.Get("user")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user,
		"exp": time.Now().Add(-time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN_SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
		return
	}
	c.SetCookie("Authorization", tokenString, 3600*24*1, "", "", false, true)
}
