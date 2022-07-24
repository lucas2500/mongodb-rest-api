package routes

import (
	"mongodb-rest-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/Product", controllers.UpsertProduct)
	app.Delete("/api/Product", controllers.DeleteProduct)
	app.Get("/api/Product", controllers.FindProducts)
}
