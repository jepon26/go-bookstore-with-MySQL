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
