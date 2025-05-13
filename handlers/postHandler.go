package handlers

import "github.com/gin-gonic/gin"

func Post(ctx *gin.Context)  {
	postTitle := ctx.Param("title")

	// Load HTML file

	// Pass content to template
}