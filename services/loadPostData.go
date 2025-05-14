package services

import (
	"encoding/json"
	"os"

	"github.com/gabrielg2020/blog/logger"
	"github.com/gabrielg2020/blog/models"
)

// Load data in `data/blogs.json`
func LoadPostData() ([]models.BlogPost, error) {
	var posts []models.BlogPost

	// Read the JSON file
	jsonFile, err := os.ReadFile("data/blogs.json")
	if err != nil {
		logger.Error("Error reading JSON file: ", err)
		return nil, err
	}

	// Unmarshall the JSON into map
	err = json.Unmarshal(jsonFile, &posts)
	if err != nil {
		logger.Error("Error unmarshaling JSON: ", err)
		return nil, err
	}

	return posts, nil
}
