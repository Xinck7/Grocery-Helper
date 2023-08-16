package data

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryStoreApi(c *gin.Context) {
	// if krogers do logic
	url := "https://api.kroger.com/v1/products?filter.brand={{BRAND}}&filter.term={{TERM}}&filter.locationId={{LOCATION_ID}}"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer {{TOKEN}}")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	// if sams do logic
}

func CalculateListPrice(c *gin.Context) {
	// input list as parameter
	//
}
