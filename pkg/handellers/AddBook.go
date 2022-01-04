package handellers

import (
	"encoding/json"
	"github.com/iamtakdir/crud-go/pkg/mock"
	"github.com/iamtakdir/crud-go/pkg/models"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	//read the request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	//append to the book mocks
	book.Id = rand.Intn(100)
	mock.Books = append(mock.Books, book)
	//send it response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")

}
