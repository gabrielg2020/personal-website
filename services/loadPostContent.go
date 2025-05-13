package services

import (
	"fmt"
	"html/template"
	"os"

	"github.com/gabrielg2020/blog/logger"
)

// Load data in `content/<title>`
func LoadPostContent(title string) (template.HTML, error) {
	// Read the HTML file
	htmlFile, err := os.ReadFile(fmt.Sprintf("content/%s.html", title))
	if err != nil {
		logger.Error("Error reading HTML file: ", err)
		return "", err
	}

	// Convert to template.HTML
	return template.HTML(htmlFile), nil
}