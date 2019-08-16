package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zgltao/jerry/controller"
	"github.com/zgltao/jerry/libs/erro"
	"github.com/zgltao/jerry/middleware"
	"github.com/zgltao/jerry/router/api/v1"
	"net/http"
)

func Load(app *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// handle no route
	app.NoRoute(func(c *gin.Context) {
		s := erro.NoRouteMatched.SetUrl(c.Request.URL.String())
		c.JSON(http.StatusNotFound, s)
	})

	// handle no method
	app.NoMethod(func(c *gin.Context) {
		s := erro.NoMethodMatched.SetUrl(c.Request.URL.String())
		c.JSON(http.StatusForbidden, s)
	})

	// apply middleware
	app.Use(middleware.CORS)
	app.Use(middleware.NoCache)
	app.Use(middleware.Secure)
	app.Use(middleware.ErrorHandler)

	app.Use(mw...)

	// mount routes
	// Example for binding JSON ({"user": "manu", "password": "123"})
	user := app.Group("/user")
	user.GET("/", middleware.LoginRequired, controller.GetUsers)
	user.POST("/login", controller.Login)
	user.POST("/register", controller.Register)

	apiv1 := app.Group("/api/v1")
	apiv1.GET("/bdcdjzmqk", v1.GetBdcdjzmqk)

	return app
}
