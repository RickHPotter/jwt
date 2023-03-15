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
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body.",
		})

		return
	}

	// Hash the pass
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to hash the password.",
		})
		return
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(hash)}

	result := initialisers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create",
		})
	}

	// Respond
	c.JSON(http.StatusAccepted, gin.H{
		"Message": "Hell Yeah, Mate!",
	})
}

func Login(c *gin.Context) {
	// Get the email and pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body.",
		})
		return
	}

	// Look up requested user
	var user models.User

	initialisers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to find user with this email.",
		})
		return
	}

	// Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to match password with this email.",
		})
		return
	}

	// Generate a JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to generate a JWT.",
		})
		return
	}

	// Send it back

	// cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorisation", tokenStr, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusAccepted, gin.H{
		"Message": tokenStr,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"Message": user.(models.User).Email + " is logged in!",
	})
}
