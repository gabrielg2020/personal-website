package routes

import "github.com/gin-gonic/gin"

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

	return router
}