package handlers

import (
	"github.com/airbornharsh/stackoverflow-backend/internal/database"
	"github.com/airbornharsh/stackoverflow-backend/pkg/helpers"
	"github.com/airbornharsh/stackoverflow-backend/pkg/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c *gin.Context) {
	var User models.User

	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Request",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	User.Password = string(hashedPassword)

	database.DB.Create(&User)

	token := helpers.GenerateToken(&User)

	c.JSON(200, gin.H{
		"message": "Registered",
		"token":   token,
	})
}

func LoginHandler(c *gin.Context) {
	var User models.User

	c.ShouldBindJSON(&User)

	var TempUser models.User

	database.DB.First(&User, "email_id = ?", User.EmailId).Scan(&TempUser)

	if TempUser.ID == 0 {
		c.JSON(404, gin.H{
			"message": "User Not Found",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(TempUser.Password), []byte(User.Password)); err != nil {
		c.JSON(401, gin.H{
			"message": "Invalid Credentials",
		})
		return
	}

	token := helpers.GenerateToken(&TempUser)

	c.JSON(200, gin.H{
		"message": "Logged In",
		"token":   token,
	})
}
