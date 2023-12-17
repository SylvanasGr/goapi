package models

import (
	"github.com/SylvanasGr/goapi/internal/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

//same us func Create(b *Book) *Book
func (b *Book) Create() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAll() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?",Id).Find(&getBook)
	return &getBook, db
}

func DeleteById(Id int64) Book{
	var book Book
	db.Where("ID=?",Id).Delete(book)
	return book
}