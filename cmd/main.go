package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/iamtakdir/crud-go/pkg/handellers"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/books", handellers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", handellers.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", handellers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handellers.UpdateBook).Methods(http.MethodPut)
	http.ListenAndServe(":4000", router)

	db, err := sql.Open("mysql", "rootuser:pass1234@/posts")
	if err != nil {
		log.Fatal("Connection has not been succeed")
	}
	// close database after all work is done
	defer db.Close()
}
