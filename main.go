package main

import (
	_ "BetFairL1/docs"
	"BetFairL1/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title        L1 BetFair
// @version      1.0
// @description  This is an API Documentation for L1 BetFair

// @contact.name   Amit
// @contact.email  amit.m@outlook.com

// @BasePath  /api/v1
func main() {
	app := fiber.New()
	app.Use(cors.New())

	// database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
