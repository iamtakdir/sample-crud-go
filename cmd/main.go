package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/iamtakdir/crud-go/pkg/handellers"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/books", handellers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", handellers.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", handellers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handellers.UpdateBook).Methods(http.MethodPut)
	http.ListenAndServe(":4000", router)

}
