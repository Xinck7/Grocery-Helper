package routes

import (
	"xinck/api/controllers"
	"xinck/api/middleware"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()

	const itemsBaseRoute = "/items"
	const itemsIDRoute = "/:id"
	r.POST(itemsBaseRoute, middleware.RequireAuth, controllers.CreateItem)
	r.GET(itemsBaseRoute, middleware.RequireAuth, controllers.GetAllItems)
	r.GET(itemsBaseRoute+itemsIDRoute, middleware.RequireAuth, controllers.GetItemByID)
	r.PUT(itemsBaseRoute+itemsIDRoute, middleware.RequireAuth, controllers.UpdateItem)
	r.DELETE(itemsBaseRoute+itemsIDRoute, middleware.RequireAuth, controllers.DeleteItem)

	const ingredientsBaseRoute = "/ingredients"
	const ingredientsIDRoute = "/:id"
	r.POST(ingredientsBaseRoute, middleware.RequireAuth, controllers.CreateIngredient)
	r.GET(ingredientsBaseRoute, middleware.RequireAuth, controllers.GetAllIngredients)
	r.GET(ingredientsBaseRoute+ingredientsIDRoute, middleware.RequireAuth, controllers.GetIngredientByID)
	r.PUT(ingredientsBaseRoute+ingredientsIDRoute, middleware.RequireAuth, controllers.UpdateIngredient)
	r.DELETE(ingredientsBaseRoute+ingredientsIDRoute, middleware.RequireAuth, controllers.DeleteIngredient)

	const listsBaseRoute = "/lists"
	const listIDRoute = "/:id"
	r.POST(listsBaseRoute, middleware.RequireAuth, controllers.CreateList)
	r.GET(listsBaseRoute, middleware.RequireAuth, controllers.GetAllLists)
	r.GET(listsBaseRoute+listIDRoute, middleware.RequireAuth, controllers.GetListByID)
	r.PUT(listsBaseRoute+listIDRoute, middleware.RequireAuth, controllers.UpdateList)
	r.DELETE(listsBaseRoute+listIDRoute, middleware.RequireAuth, controllers.DeleteList)

	const recipesBaseRoute = "/recipes"
	const recipesIDRoute = "/:id"
	r.POST(recipesBaseRoute, middleware.RequireAuth, controllers.CreateRecipe)
	r.GET(recipesBaseRoute, middleware.RequireAuth, controllers.GetAllRecipes)
	r.GET(recipesBaseRoute+recipesIDRoute, middleware.RequireAuth, controllers.GetRecipeByID)
	r.PUT(recipesBaseRoute+recipesIDRoute, middleware.RequireAuth, controllers.UpdateRecipe)
	r.DELETE(recipesBaseRoute+recipesIDRoute, middleware.RequireAuth, controllers.DeleteRecipe)

	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	r.GET("/validate", middleware.RequireAuth, controllers.ValidateUserSession)
	r.GET("/logout", controllers.LogoutUser)

	r.Run()
}
