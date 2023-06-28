package controllers

import (
	"github.com/feebrian/gotodolist/app/entities"
	"github.com/feebrian/gotodolist/config"
)

func CheckUserById() {
	var err error
	var user entities.User

	if err = config.DB.Where("id =", user.ID).First(&user).Error; err != nil {
		panic("user not found.")
	}

}
