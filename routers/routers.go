package routers

import (
	"net/http"

	"github.com/RickHPotter/jwt/controllers"
	"github.com/RickHPotter/jwt/middleware"
	"github.com/gin-gonic/gin"
)

func LoadRoute() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "Pong",
		})
	})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()

	http.ListenAndServe(":3500", nil)
}
