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
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var book models.Book
	//read the request body
	defer r.Body.Close()
	json.Unmarshal(body, &book)

	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")

}
