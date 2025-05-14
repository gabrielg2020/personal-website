package handlers

import (
	"net/http"

	"github.com/gabrielg2020/blog/logger"
	"github.com/gabrielg2020/blog/services"
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	// Load posts
	posts, err := services.LoadPostData()
	if err != nil {
		logger.Error("Failed to load blog post data: ", err)
		ErrorPage(ctx)
		return
	}

	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"Posts": posts,
	})
}
