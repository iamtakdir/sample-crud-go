package handellers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/iamtakdir/crud-go/pkg/models"
	"log"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/bookstoredb")
	if err != nil {
		log.Fatal("Connection has not been succeed in Add book")
	}
	// close database after all work is done
	defer db.Close()

	var book models.Book
	//read the request body
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&book)

	insert, err := db.Prepare("INSERT INTO `books` (`id`,`title`, `author`,`desc`) VALUES (?,?,?,?) ")
	if err != nil {
		log.Fatal("error in inserting")
	}
	fmt.Println("Printing Book", book)
	fmt.Println("printing Insert", insert)
	insert.Exec(&book.Id, &book.Title, &book.Author, &book.Desc)
	if err != nil {
		log.Fatal("Error in execution ")
	}
	fmt.Println("Inserted a")
	fmt.Println("Inserted successfully")

	////append to the book mocks
	//book.Id = rand.Intn(100)
	//mock.Books = append(mock.Books, book)
	////send it response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")

}
