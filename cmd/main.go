package main

import (
	"log"

	"github.com/go-layer-architecture/api"
	"github.com/go-layer-architecture/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	utils.ConnectDB()
	api.Router(app)

	log.Fatal(app.Listen(":3000"))
}
