package repository

import (
	"github.com/go-layer-architecture/internal/domain"
	"github.com/go-layer-architecture/pkg/utils"
)

func SaveUser(u *domain.CreateUser) error {
	user := domain.User{
		Name:     u.Name,
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
