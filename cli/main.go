package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

type BlogPost struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
}

func main() {
	// Check if argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Error: Missing required argument")
		fmt.Println("Usage:", os.Args[0], "<blogTitle>")
		os.Exit(1)
	}

	// Define command line flag
	flag.Parse()

	blogTitle := os.Args[1]

	createNewBlog(blogTitle)	
}

func createNewBlog(blogTitle string) {
	// Create new HTML file with correct name
	file, err := os.Create(fmt.Sprintf("../content/%s.html", blogTitle))
	if err != nil {
		fmt.Println("Error creating file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	// Write inital content
	initialHtmlContent := "<!DOCTYPE html>"

	_, err = file.WriteString(initialHtmlContent)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		os.Exit(1)
	}

	// Read existing JSON file
	jsonData, err := os.ReadFile("../data/blogs.json")
	if err != nil {
		fmt.Println("Error reading JSON file: ", err)
		os.Exit(1)
	}

	// Parse existing posts
	var posts []BlogPost
	err = json.Unmarshal(jsonData, &posts)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
		os.Exit(1)
	}

	// Create a new entry
	currentDate := time.Now().Format("2 January, 2006")
	newPost := BlogPost {
		Title: blogTitle,
		Date: currentDate,
	}

	// Convert back to JSON
	posts = append(posts, newPost)
	updateJSON, err := json.MarshalIndent(posts, "", " ")
	if err != nil {
		 fmt.Println("Error creating JSON: ", err)
		 os.Exit(1)
	}

	// Write back to file
	err = os.WriteFile("../data/blogs.json", updateJSON, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully created blog!")
}