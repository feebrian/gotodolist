package handlers

import (
	"github.com/go-layer-architecture/internal/app/services"
	"github.com/go-layer-architecture/internal/domain"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user domain.CreateUser
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
			"user":    nil,
		})
	}

	if err := services.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
			"user":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully create User",
		"user":    user,
	})
}

func FindUsers(c *fiber.Ctx) error {
	var users []domain.User

	if err := services.FindUsers(&users); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
			"user":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Users found",
		"users":   users,
	})
}
