package handellers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/iamtakdir/crud-go/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {

	//	Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//Read the body

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	//unmarshal from body and update new object

	var updateBook models.Book
	json.Unmarshal(body, &updateBook)

	//Find old book and replace his data with new data

	var book models.Book

	//find book from database

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	//replace data with new one

	book.Title = updateBook.Title
	book.Author = updateBook.Author
	book.Desc = updateBook.Desc
	//save data
	h.DB.Save(&book)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Updated")

}
