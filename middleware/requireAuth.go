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

func Abort(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"Error": errorMessage,
	})
	c.AbortWithStatus(http.StatusUnauthorized)
}

func RequireAuth(c *gin.Context) {
	// Get the cookie off req
	tokenStr, err := c.Cookie("Authorisation")

	if err != nil {
		Abort(c, "Cookie nonexistent or not valid.")
	}

	// Decode
	// Parse takes the token string and a function for looking up the key
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// method of signing has to be the same, a bit lost here though
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		Abort(c, "Failed to decode cookie.")
	}

	// Validation
	// method of claims also has to match, I suppose,
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			Abort(c, "Session Cookie has expired.")
		}
		// Find the user with Token Sub
		var user models.User
		// the sub is the subject claim of the JWT, always case-sensitive, and optional
		initialisers.DB.First(&user, claims["sub"])

		// user.ID will be zero in case DB.First failed to find a user
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to request body so that it can be used in the controllers
		c.Set("user", user)

		// this is a handle before another handle, therefore you need to
		// tell the context to proceed to the next handle, as far as I can see
		c.Next()
	} else {
		Abort(c, "Something wrong with JWT Claims.")
	}

}
