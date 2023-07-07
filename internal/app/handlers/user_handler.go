package handlers

import (
	"strconv"

	"github.com/go-layer-architecture/internal/app/services"
	"github.com/go-layer-architecture/internal/domain"
	"github.com/go-layer-architecture/internal/repository"
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

func GetUser(c *fiber.Ctx) error {

	// get param
	userID := c.Params("id")
	uid, _ := strconv.Atoi(userID)

	user, err := repository.FindUserById(uint(uid))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
			"user":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User found",
		"users":   user,
	})

}
