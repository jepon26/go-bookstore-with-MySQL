package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ahkil/go-bookstore/pkg/models"
	"github.com/ahkil/go-bookstore/pkg/utils"
	"github.com/akhil/go-bookstore/pkg/models"
	"github.com/gorilla/mux"
)

var NewBook models.Book


func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks:= models.GetAllBooks()
	res, _:= json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkgaplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
}
bookDetails, _:= models.GetBookById(ID)
res, _:= json.Marshal(bookDetails)
w.Header().Set("Content-Type", "pkgaplication/json")
w.WriteHeader(http.StatusOK)
w.Write(res)
}


func CreateBook(w http.ResponseWriter, r *http.Request){
CreateBook := models.Book{}
utils.parseBody(r, CreateBook)
b:= CreateBook.CreateBook()
res, _:= json.Marshal(b)
w.WriteHeader(http.StatusOK)
w.Write(res)
}


