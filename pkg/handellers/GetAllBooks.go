package handellers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/iamtakdir/crud-go/pkg/models"
	"log"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bookstoredb")
	if err != nil {
		log.Fatal("Connection has not been succeed")
	}
	// close database after all work is done
	defer db.Close()

	// query all data
	rows, err := db.Query("select * from books")
	fmt.Println(rows)
	if err != nil {
		log.Fatal("Error in fetching all data")
	}

	var book models.Book

	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Desc)

		fmt.Println(book)
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
}
