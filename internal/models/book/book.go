package book

import (
	"example/bookAPI/internal/models/author"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	AuthorID int `json:"authorId"`
	Author   author.Author
	Name     string `json:"name"`
}
