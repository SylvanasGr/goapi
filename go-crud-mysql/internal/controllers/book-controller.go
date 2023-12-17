package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SylvanasGr/goapi/internal/models"
	"github.com/SylvanasGr/goapi/internal/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r* http.Request){
	newBooks := models.GetAll()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r* http.Request){
	vars := mux.Vars(r)
	bookId := vars["id"]
	fmt.Printf("bookId is %v\n",bookId)
	id, err := strconv.ParseInt(bookId,0,0) 

	if err != nil {
		fmt.Printf("Error is %v",err)
	}

	bookDetails, _ := models.GetById(id)
	fmt.Printf("Id for book is %v\n",id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter,r* http.Request){
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.Create()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter,r* http.Request){
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r* http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r,updateBook)
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetById(id)

	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Name
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	res , _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}