package models

import "html/template"

type BlogPost struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Content template.HTML
}
