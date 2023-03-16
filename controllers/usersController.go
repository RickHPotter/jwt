package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/RickHPotter/jwt/initialisers"
	"github.com/RickHPotter/jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get the email/pass off req body

	if c.Bind(&(models.Body)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body.",
		})

		return
	}

	// Hash the pass
	hash, err := bcrypt.GenerateFromPassword([]byte(models.Body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to hash the password.",
		})
		return
	}

	// Create the user
	user := models.User{Email: models.Body.Email, Password: string(hash)}

	result := initialisers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create",
		})
	}

	// Respond
	c.JSON(http.StatusAccepted, gin.H{
		"Message": "Hell Yeah, Mate! " + user.Email + " is signed up.",
	})
}

func Login(c *gin.Context) {
	// Get the email and pass off req body

	if c.Bind(&(models.Body)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body.",
		})
		return
	}

	// Look up requested user
	var user models.User

	initialisers.DB.First(&user, "email = ?", models.Body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to find user with this email.",
		})
		return
	}

	// Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(models.Body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to match password with this email.",
		})
		return
	}

	// Generate a JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to generate a JWT.",
		})
		return
	}

	// Create a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorisation", tokenStr, 3600*24*30, "", "", false, true)

	// Send JWT back
	c.JSON(http.StatusAccepted, gin.H{
		"Message": tokenStr,
	})
}

// mock func just to see middleware working
func Validate(c *gin.Context) {

	// this is only possible because there was a c.Set("user", user) in requireAuth
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "This is not supposed to happen, by the way.",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": user.(models.User).Email + " is logged in!",
	})
}
