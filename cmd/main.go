package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/iamtakdir/crud-go/pkg/db"
	"github.com/iamtakdir/crud-go/pkg/handellers"
	"net/http"
)

func main() {

	DB := db.Init()
	h := handellers.New(DB)

	router := mux.NewRouter()

	//Routes

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)
	http.ListenAndServe("localhost:4000", router)

}
