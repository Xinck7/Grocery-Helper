package routes

import (
	"xinck/api/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	route := gin.Default()

	const itemsBaseRoute = "/items"
	const itemsIDRoute = "/:id"
	route.POST(itemsBaseRoute, controllers.CreateItem)
	route.GET(itemsBaseRoute, controllers.GetAllItems)
	route.GET(itemsBaseRoute+itemsIDRoute, controllers.GetItemByID)
	route.PUT(itemsBaseRoute+itemsIDRoute, controllers.UpdateItem)
	route.DELETE(itemsBaseRoute+itemsIDRoute, controllers.DeleteItem)

	const ingredientsBaseRoute = "/ingredients"
	const ingredientsIDRoute = "/:id"
	route.POST(ingredientsBaseRoute, controllers.CreateIngredient)
	route.GET(ingredientsBaseRoute, controllers.GetAllIngredients)
	route.GET(ingredientsBaseRoute+ingredientsIDRoute, controllers.GetIngredientByID)
	route.PUT(ingredientsBaseRoute+ingredientsIDRoute, controllers.UpdateIngredient)
	route.DELETE(ingredientsBaseRoute+ingredientsIDRoute, controllers.DeleteIngredient)

	const listsBaseRoute = "/lists"
	const listIDRoute = "/:id"
	route.POST(listsBaseRoute, controllers.CreateList)
	route.GET(listsBaseRoute, controllers.GetAllLists)
	route.GET(listsBaseRoute+listIDRoute, controllers.GetListByID)
	route.PUT(listsBaseRoute+listIDRoute, controllers.UpdateList)
	route.DELETE(listsBaseRoute+listIDRoute, controllers.DeleteList)

	const recipesBaseRoute = "/recipes"
	const recipesIDRoute = "/:id"
	route.POST(recipesBaseRoute, controllers.CreateRecipe)
	route.GET(recipesBaseRoute, controllers.GetAllRecipes)
	route.GET(recipesBaseRoute+recipesIDRoute, controllers.GetRecipeByID)
	route.PUT(recipesBaseRoute+recipesIDRoute, controllers.UpdateRecipe)
	route.DELETE(recipesBaseRoute+recipesIDRoute, controllers.DeleteRecipe)

	route.Run()
}
