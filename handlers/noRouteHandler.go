package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoRoute(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "noRoute.html", gin.H{})
}
