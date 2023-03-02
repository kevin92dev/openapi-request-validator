package router

import (
	"github.com/gin-gonic/gin"
	"main/api/handlers"
	"main/internal/infrastructure/middleware"
)

func GenerateRouting(sentryHandler gin.HandlerFunc) *gin.Engine {
	app := gin.Default()
    
	app.Use(middleware.ValidateRequest())

	router := app.Group("/")
	handlers.Routes(router)
    
	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Route not found"})
	})

	return app
}
