package handellers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/iamtakdir/crud-go/pkg/models"
	"net/http"
	"strconv"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {

	//Read unique id from url

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	var book models.Book
	// find a book by id and return it.
	if result := h.DB.Find(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	
	//encode json
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)

}
