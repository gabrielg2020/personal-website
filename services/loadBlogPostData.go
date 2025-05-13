package services

import (
	"encoding/json"
	"os"

	"github.com/gabrielg2020/blog/logger"
	"github.com/gabrielg2020/blog/models"
)

// Load data in `./data/blogs.json` into a map of BlogPost structs
func LoadBlogPostData() (map[string][]models.BlogPost, error) {
	var years map[string][]models.BlogPost

	// Read the JSON file
	jsonFile, err := os.ReadFile("data/blogs.json")
	if err != nil {
		logger.Error("Error opening JSON file: ", err)
		return nil, err
	}

	// Unmarshall the JSON into map
	err = json.Unmarshal(jsonFile, &years)
	if err != nil {
		logger.Error("Error unmarshaling JSON: ", err)
		return nil, err
	}

	return years, nil
}