package server

import (
	"github.com/alexohneander/codesharego/api"
	"github.com/gofiber/fiber/v2"
)

func Listen() {
	app := fiber.New(fiber.Config{
		ServerHeader:          "Fiber",
		AppName:               "CodeshareGo v0.1.0",
		DisableStartupMessage: true,
	})

	// API
	apiGroup := app.Group("/api") // /api

	v1 := apiGroup.Group("/v1")     // /api/v1
	v1.Get("/rooms", api.ListRooms) // /api/v1/rooms

	// Prepare a static middleware to serve the built React files.
	app.Static("/", "./web/build")

	// Prepare a fallback route to always serve the 'index.html', had there not be any matching routes.
	app.Static("*", "./web/build/index.html")

	app.Listen(":4000")
}
