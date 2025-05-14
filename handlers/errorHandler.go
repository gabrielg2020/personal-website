package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorPage(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "error.html", gin.H{})
}
