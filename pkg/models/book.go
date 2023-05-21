package models

import (
	"github.com/jinzhu/gorm"
	"github.com/akhil/go-bookstore/pkg/config"
)

// DB - DB Instance
var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Author      string `gorm:"column:author" json:"author"`
	Publication string `gorm:"column:publication" json:"publication"`
}


func init (){
	config.Connect()
	db = config.GetDB()
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db:=db.Where("ID=" , Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}

