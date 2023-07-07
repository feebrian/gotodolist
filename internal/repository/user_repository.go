package repository

import (
	"github.com/go-layer-architecture/internal/domain"
	"github.com/go-layer-architecture/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SaveUser(u *domain.CreateUser) error {
	user := domain.User{
		Username: u.Username,
		Email:    u.Email,
		Password: utils.HashPassword(u.Password),
	}
	result := utils.DB.Create(&user)
	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func FindUsers(u *[]domain.User) error {
	return utils.DB.Find(u).Error
}

func FindUserById(uid uint) (domain.User, error) {
	var user domain.User
	if err := utils.DB.Where("id = ?", uid).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return user, fiber.NewError(fiber.StatusNotFound, "user not found")
		default:
			return user, fiber.NewError(fiber.StatusInternalServerError, "failed find user")
		}
	}

	return user, nil
}
