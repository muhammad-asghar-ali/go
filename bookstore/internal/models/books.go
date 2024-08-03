package models

import (
	"bookstore/internal/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type (
	Books struct {
		gorm.Model
		Name        string `gorm:"" json:"name"`
		Author      string `gorm:"" json:"author"`
		Publication string `gorm:"" json:"publication"`
	}
)

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Books{})
}

func (b *Books) CreateBook() *Books {
	db.Create(&b)
	return b
}

func (b *Books) GetBooks() []*Books {
	books := make([]*Books, 0)
	db.Find(books)
	return books
}

func (b *Books) GetBookByID(id int64) (*Books, *gorm.DB) {
	book := Books{}
	db := db.Where("ID=?", id).Find(&book)
	return &book, db
}

func (b *Books) DeleteBook(id int64) *Books {
	book := Books{}
	db.Where("ID=?", id).Delete(&book)
	return &book
}
