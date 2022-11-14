package models

import (
	"errors"

	"github.com/gastighost/book-management-system/pkg/configs"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	configs.Connect()
	db = configs.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) Create() (*Book, error) {
	err := db.Create(&b).Error

	if err != nil {
		return nil, errors.New("there was an error creating your book")
	}

	return b, nil
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Find(&book)
	db.Delete(&book)
	return book
}
