package api

import (
	"github.com/go-layer-architecture/internal/app/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Router(app *fiber.App) {
	app.Use(cors.New(cors.ConfigDefault))

	app.Post("/user", handlers.CreateUser)
	app.Get("/users", handlers.FindUsers)
	app.Get("/users/:id", handlers.GetUser)
}
