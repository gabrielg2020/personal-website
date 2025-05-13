package models

type BlogPost struct {
	Title string `json:"title"`
	ShortDescription string `json:"shortDescription"`
	Tags []string `json:"tags"`
}