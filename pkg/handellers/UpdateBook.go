package handellers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/iamtakdir/crud-go/pkg/models"
	"log"
	"net/http"
	"strconv"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bookstoredb")
	if err != nil {
		log.Fatal("Connection has not been succeed in Add book")
	}
	//	Read dynamic id parameter

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	fmt.Println(id)
	//	Read request body
	defer r.Body.Close()
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	update, err := db.Prepare("UPDATE books SET title = ?,author = ?,desc= ? WHERE id = ?")
	fmt.Println("Printing Update", update)
	if err != nil {
		log.Fatal("Error in updating")
	}

	update.Exec(&book.Title, &book.Author, &book.Desc, id)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Updated")

}
