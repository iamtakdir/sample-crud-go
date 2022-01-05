package handellers

import (
	"encoding/json"
	"fmt"
	"github.com/iamtakdir/crud-go/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {

	// read data from body
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	var book models.Book
	// unmarshal data and bind it with book

	json.Unmarshal(body, &book)
	//Create data 
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")

}
