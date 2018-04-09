package main

import (
	"github.com/dmirubtsov/swag-example/controllers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/dmirubtsov/swag-example/docs"
)

var router = gin.Default()

func SetupRouter() *gin.Engine {
	v1 := router.Group("v1")
	// Documentation
	router.GET("/", docsRedirect)
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// User
	{
		v1.PUT("/user", controllers.User.Update)
	}

	return router
}

func docsRedirect(c *gin.Context) {
	c.Redirect(301, "/v1/swagger/index.html")
}
