package models

import (
	"fmt"

	"github.com/LuisSilva7/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	var existingBook Book
	if err := db.Where("name=?", b.Name).First(&existingBook).Error; err == nil {
		return nil, fmt.Errorf("book already exists")
	}
	db.Create(&b)
	return b, nil
}

func GetAllBooks() ([]Book, error) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookByID(ID int64) (*Book, error) {
	var book Book
	if err := db.Where("ID=?", ID).Find(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (b *Book) UpdateBook() (*Book, error) {
	db.Save(&b)
	return b, nil
}

func DeleteBook(ID int64) (*Book, error) {
	var book Book
	if err := db.Where("ID=?", ID).Delete(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
