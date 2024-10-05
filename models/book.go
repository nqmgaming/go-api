package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
	ISBN          string `json:"isbn"`
}
