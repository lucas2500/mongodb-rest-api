package main

import (
	"log"
	"mongodb-rest-api/database"
	"mongodb-rest-api/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("There was an error when trying to load .env file!!", err)
	}

	database.MongoURI = os.Getenv("MONGO_URI")
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

	var res = make(map[string]string)

	res["Ping"] = "Pong"

	return c.JSON(res)
}
