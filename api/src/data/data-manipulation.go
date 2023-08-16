package data

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryStoreApi(c *gin.Context) {
	//? if krogers do logic
	//* from Krogers API -> adapted ioutil.ReadAll to io.ReadAll
	url := "https://api.kroger.com/v1/products?filter.brand={{BRAND}}&filter.term={{TERM}}&filter.locationId={{LOCATION_ID}}"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer {{TOKEN}}")

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

func CalculateListPrice(c *gin.Context) {
	//? input list as parameter
	//? Create array of all the price values
	//? Sum the elements in the array
	//? return the sum
	//? call controllers.UpdateList() with the new value
	//? ensure only updates the price and not entering null values on top
}

func CalculateRecipePrice(c *gin.Context) {
	//? input recipe as parameter
	//? Create array of all the price values
	//? Sum the elements in the array
	//? return the sum
	//? call controllers.UpdateRecipe() with the new value
	//? ensure only updates the price and not entering null values on top
}

func UpdateAllItemInfo(c *gin.Context) {
	//? for each in batches/long running task 1x a day run the item update on all database items
}

func UpdateAllIngredientInfo(c *gin.Context) {
	//? for each in batches/long running task 1x a day run the ingredient update on all database ingredient
}

func UpdateAllListInfo(c *gin.Context) {
	//? for each in batches/long running task 1x a day run the list update on all database list
}

func UpdateAllRecipeInfo(c *gin.Context) {
	//? for each in batches/long running task 1x a day run the recipe update on all database recipe
}
