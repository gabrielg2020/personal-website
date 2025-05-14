package routes

import (
	"github.com/gabrielg2020/blog/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// API routes
	api := router.Group("/api")
	{
		api.GET("/post/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.JSON(200, gin.H{
				"name": name,
			})
		})
	}

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Load static files
	router.Static("/css", "./static/styles")

	// Page routes
	router.GET("/", handlers.Home)
	post := router.Group("/post")
	{
		post.GET("/:title", handlers.Post)
	}

	// No route
	router.NoRoute(handlers.NoRoute)

	return router
}
