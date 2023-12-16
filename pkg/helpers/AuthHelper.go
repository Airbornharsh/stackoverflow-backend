package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/airbornharsh/stackoverflow-backend/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(user *models.User) string {
	JWT_SECRET := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"userId":   user.ID,
		"name":     user.DisplayName,
		"username": user.Username,
		"emailId":  user.EmailId,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString([]byte(JWT_SECRET))

	return tokenString
}

func GetClaims(c *gin.Context, tokenString string) (models.User, error) {
	JWT_SECRET := os.Getenv("JWT_SECRET")
	var user models.User

	tokenString = tokenString[7:]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return user, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return user, fmt.Errorf("invalid claims")
	}

	user.ID = uint(claims["userId"].(float64))
	user.DisplayName = claims["name"].(string)
	user.EmailId = claims["emailId"].(string)
	user.Username = claims["username"].(string)

	return user, nil
}
