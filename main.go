package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/subosito/gotenv"
	"go.mod/src/config"
	"go.mod/src/helper"
	"go.mod/src/routes"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helper.Migrate()
	// defer config.DB.Close()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))

	routes.Router(app)
	app.Listen(":8080")

}
