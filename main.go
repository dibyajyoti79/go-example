package main

import (
	"go-example/handlers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Get("/", handlers.Hello)
	app.Get("/user/:id", handlers.GetUser)
	app.Post("/user", handlers.CreateUser)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
