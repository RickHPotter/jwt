package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/RickHPotter/jwt/initialisers"
	"github.com/RickHPotter/jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie off req
	tokenStr, err := c.Cookie("Authorisation")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode / validate it
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	// closure
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// Find the user iwth Token Sub
		var user models.User
		initialisers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to req
		c.Set("user", user)

		// Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
