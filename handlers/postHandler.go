package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gabrielg2020/blog/logger"
	"github.com/gabrielg2020/blog/models"
	"github.com/gabrielg2020/blog/services"
	"github.com/gin-gonic/gin"
)

func Post(ctx *gin.Context) {
	title := ctx.Param("title")

	// Check if file exists
	_, err := os.Stat(fmt.Sprintf("content/%s.html", title))
	if err != nil {
		if os.IsNotExist(err) {
			NoRoute(ctx) // Move to no route handler if post does not exist
			return
		}
	}

	// Load HTML file
	content, err := services.LoadPostContent(title)
	if err != nil {
		logger.Error("Error loading post content: ", err)
		return
	}

	// Load posts
	posts, err := services.LoadPostData()
	if err != nil {
		logger.Error("Error loading post data: ", err)
		return
	}

	// Find the post by title
	var postToShow *models.BlogPost
	for _, post := range posts {
		if post.Title == title {
			post.Content = content
			postToShow = &post
			break // Exit loop once the post is found
		}
	}

	// Pass content to template
	if postToShow != nil {
		ctx.HTML(http.StatusOK, "post.html", gin.H{
			"Title":   postToShow.Title,
			"Date":    postToShow.Date,
			"Content": postToShow.Content,
		})
	} else {
		logger.Error("Post not found: ", title)
	}
}
