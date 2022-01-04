package handellers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/iamtakdir/crud-go/pkg/mock"
	"github.com/iamtakdir/crud-go/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	//	Read dynamic id parameter

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	//	Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)

	}

	var updateBook models.Book

	json.Unmarshal(body, &updateBook)

	//	iterate over all books

	for index, book := range mock.Books {
		if book.Id == id {
			book.Title = updateBook.Title
			book.Author = updateBook.Author
			book.Desc = updateBook.Desc

			mock.Books[index] = book

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
	//	update and send response
}
