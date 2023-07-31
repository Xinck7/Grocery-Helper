package routes

import (
	"xinck/api/src/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	route := gin.Default()

	// items
	route.POST("/items", controllers.CreateItem)
	route.GET("/items", controllers.GetAllItems)
	route.PUT("/items/:idTodo", controllers.UpdateItem)
	route.DELETE("/items/:idTodo", controllers.DeleteItem)

	// ingredients
	route.POST("/ingredients", controllers.CreateIngredient)
	route.GET("/ingredients", controllers.GetAllIngredients)
	route.PUT("/ingredients/:idIngredient", controllers.UpdateIngredient)
	route.DELETE("/ingredients/:idIngredient", controllers.DeleteIngredient)

	// recipes
	route.POST("/recipes", controllers.CreateRecipe)
	route.GET("/recipes", controllers.GetAllRecipes)
	route.PUT("/recipes/:idRecipe", controllers.UpdateRecipe)
	route.DELETE("/recipes/:idRecipe", controllers.DeleteRecipe)

	// lists
	route.POST("/lists", controllers.CreateList)
	route.GET("/lists", controllers.GetAllLists)
	route.PUT("/lists/:idlist", controllers.UpdateList)
	route.DELETE("/lists/:idList", controllers.DeleteList)

	// Run
	route.Run()
}
