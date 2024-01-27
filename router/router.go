package router

import (
	"github.com/gorilla/mux"
	"github.com/rishavqwerty7/BookwormApi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/books/allBooks", controller.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/create", controller.CreateBook).Methods("POST")
	router.HandleFunc("/books/update/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/delete/{id}", controller.DeleteOneBook).Methods("DELETE")
	router.HandleFunc("/books/deleteAll", controller.DeleteAllBooks).Methods("DELETE")

	return router
}
