package authcontrollers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/feebrian/gotodolist/app/entities"
	"github.com/feebrian/gotodolist/config"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	// to register user if user doesn't have an account
	var err error

	var input entities.RegisterUser
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// make password hashing
	hashPassword := sha256.Sum256([]byte(input.Password))
	hashPasswordStr := hex.EncodeToString(hashPassword[:])

	user := entities.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashPasswordStr,
	}

	config.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"message": "Sucessfully register, please login now."})
}

func Login(c *gin.Context) {
	// to login user
}

func Logout(c *gin.Context) {
	// to logout user
}
