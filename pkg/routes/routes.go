package routes

import (
	"github.com/abhisenberg/go-books/pkg/controllers"

	"github.com/gorilla/mux"
)

var DeclareRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
