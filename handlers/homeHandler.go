package handlers

import (
	"net/http"

	"github.com/gabrielg2020/blog/logger"
	"github.com/gabrielg2020/blog/services"
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context)  {
	years, err := services.LoadBlogPostData()
	if err != nil {
		logger.Error("Failed to load blog post data: ", err)
	}

	logger.Info(years["2025"][0].Title)
	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"2025": years["2025"],
		"2024": years["2024"],
	})
}