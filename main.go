package main

import (
	"log"
	"mongodb-rest-api/database"
	"mongodb-rest-api/entities"
	"mongodb-rest-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("There was an error when trying to load .env file!!", err)
	}

}

func main() {

	database.InitDatabase()

	defer database.DisconnectFromMongo()

	app := fiber.New()
	app.Get("/", Hello)

	routes.SetupRoutes(app)

	app.Listen(":3000")

}

func Hello(c *fiber.Ctx) error {

	var res entities.Hello
	res.Ping = "Pong"

	return c.JSON(res)
}
