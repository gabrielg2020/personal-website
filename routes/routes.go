package routes

import (
	"net/http"

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
	router.LoadHTMLGlob("views/**/*.html")

	// Load static files
	router.Static("/css", "./static/styles")

	// Page routes
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", gin.H{})
	})

	router.GET("/about", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "about.html", gin.H{})
	})

	router.GET("/projects", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "projects.html", gin.H{})
	})

	router.GET("/reading-list", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "readingList.html", gin.H{})
	})

	router.GET("/blogs", handlers.Blog)
	post := router.Group("/blogs")
	{
		post.GET("/:title", handlers.Post)
	}

	// No route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "noRoute.html", gin.H{})
	})

	return router
}
