package api

import (
	"github.com/go-layer-architecture/internal/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Post("/user", handlers.CreateUser)
	app.Get("/users", handlers.GetUser)
}
