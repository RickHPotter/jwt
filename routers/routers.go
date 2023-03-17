package routers

import (
	"net/http"
	"os"

	"github.com/RickHPotter/jwt/controllers"
	"github.com/RickHPotter/jwt/middleware"
	"github.com/gin-gonic/gin"
)

func LoadRoute() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	// [1]
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()

	http.ListenAndServe(os.Getenv("PORT"), nil)
}

// ! [1]
// ! Multiple handlers can be passed as parameters of a r.Method()
// ! They are processed in that order making middleware able to be run before or after
// ! In this case, for /validate to be displayed, requireAuth has to return true
// ! and by that, I mean requireAuth not calling abortWithStatus(http.StatusUnauthorized).
