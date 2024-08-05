package models

import (
	"bookstore/internal/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type (
	Book struct {
		gorm.Model
		Name        string `gorm:"" json:"Name"`
		Author      string `gorm:"" json:"Author"`
		Publication string `gorm:"" json:"Publication"`
	}
)

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func (b *Book) GetBooks() []*Book {
	books := make([]*Book, 0)
	db.Find(&books)
	return books
}

func (b *Book) GetBookByID(id int64) (*Book, *gorm.DB) {
	book := Book{}
	db := db.Where("ID=?", id).Find(&book)
	return &book, db
}

func (b *Book) DeleteBook(id int64) *Book {
	book := Book{}
	db.Where("ID=?", id).Delete(&book)
	return &book
}

func (b *Book) UpdateBook() *Book {
	db.Save(&b)
	return b
}
