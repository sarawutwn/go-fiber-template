package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sarawutwn/gofiber-template/config"
	"github.com/sarawutwn/gofiber-template/database"
	"github.com/sarawutwn/gofiber-template/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(config.CorsConfigDefault))
	database.ConnectDB()
	router.SetupRoutes(app)
	app.Listen(":6000")
}
