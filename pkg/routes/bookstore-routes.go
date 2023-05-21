package routes

import (
	"github.com/gorilla/mux"
	"github.com/akhil/go-bookstore/pkg/controllers"
)


const bookEndpoint = "/book/{bookId}"


var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc(bookEndpoint, controllers.GetBookById).Methods("GET")
	router.HandleFunc(bookEndpoint, controllers.UpdateBook).Methods("PUT")
	router.HandleFunc(bookEndpoint, controllers.DeleteBook).Methods("DELETE")
}



