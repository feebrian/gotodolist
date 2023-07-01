package authcontrollers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/feebrian/gotodolist/app/entities"
	"github.com/feebrian/gotodolist/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	result := config.DB.Create(&user)
	if result.RowsAffected == 1 {
		c.JSON(http.StatusCreated, gin.H{"message": "Sucessfully register, please login now."})
		return
	}

	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": result.Error})
}

func Login(c *gin.Context) {
	// to login user
	var err error

	// get user input
	var input entities.LoginUser
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// check if user existed in database
	var user entities.User
	if err = config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password is wrong"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// check user password
	hashPassword := sha256.Sum256([]byte(input.Password))
	hashPasswordStr := hex.EncodeToString(hashPassword[:])
	// log.Fatal(hashPasswordStr)

	if err = config.DB.Model(&user).Where("password = ?", hashPasswordStr).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password is wrong"})
			return
		default:
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "berhasil login"})

}

func Logout(c *gin.Context) {
	// to logout user
}
