package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"xinck/api/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func QueryStoreApi(c *gin.Context) {
	//? if krogers do logic
	//* from Krogers API -> adapted ioutil.ReadAll to io.ReadAll
	//Kroger api token load from .env
	err := godotenv.Load(".env")
	if err != nil {
		panic("error loading .env file")
	}
	var token = os.Getenv("token")
	url := "https://api.kroger.com/v1/products?filter.brand={{BRAND}}&filter.term={{TERM}}&filter.locationId={{LOCATION_ID}}"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer"+token)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	//? if sams do logic
}

func CalculateItemPrice(c *gin.Context) {
	//? input item as parameter
	//? call query store api method with the name of the item
	//? call controllers.UpdateItem() with new value
}

func CalculateIngredientPrice(c *gin.Context) {
	//? input ingredient as parameter
	//? call query store api method with the name of the item
	//? call controllers.UpdateIngredient() with new value
}

func CalculateListPrice(list *models.List) int64 {
	var totalPrice int64

	for _, item := range list.Items {
		totalPrice += item.Price * int64(item.Quantity)
	}

	for _, ingredient := range list.Ingredients {
		totalPrice += ingredient.Price * int64(ingredient.Quantity)
	}

	for _, recipe := range list.Recipes {
		for _, ingredient := range recipe.Ingredients {
			totalPrice += ingredient.Price * int64(ingredient.Quantity)
		}
	}
	return totalPrice
}

func CalculateRecipePrice(recipe *models.Recipe) int64 {
	var totalPrice int64

	for _, ingredient := range recipe.Ingredients {
		totalPrice += ingredient.Price * int64(ingredient.Quantity)
	}

	return totalPrice
}

func UpdateAllItemInfo(c *gin.Context) {
	//? Make it a go routine
	//? for each in batches/long running task 1x a day run the item update on all database items
}

func UpdateAllIngredientInfo(c *gin.Context) {
	//? Make it a go routine
	//? for each in batches/long running task 1x a day run the ingredient update on all database ingredient
}

func UpdateAllListInfo(c *gin.Context) {
	//? Make it a go routine
	//? for each in batches/long running task 1x a day run the list update on all database list
}

func UpdateAllRecipeInfo(c *gin.Context) {
	//? for each in batches/long running task 1x a day run the recipe update on all database recipe
}
