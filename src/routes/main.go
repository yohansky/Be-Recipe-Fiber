package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mod/src/controller"
	"go.mod/src/middleware"
)

func Router(app *fiber.App) {
	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)

	app.Use(middleware.IsAuth)

	app.Get("/user", controller.User)
	app.Post("/logout", controller.Logout)

	app.Get("/users", controller.AllUsers)
	app.Get("/user/:id", controller.GetUser)
	app.Put("/user/:id", controller.UpdateUser)
	app.Delete("/user/:id", controller.DeleteUser)

	app.Get("/recipes", controller.AllRecipes)
	app.Post("/recipes", controller.CreateRecipe)
	app.Get("/recipe/:id", controller.GetRecipe)
	app.Put("/recipe/:id", controller.UpdateRecipe)
	app.Delete("/recipe/:id", controller.DeleteRecipe)

}
