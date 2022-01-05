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

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	var book models.Book
	// find a book by id
	if result := h.DB.Find(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	//Delete that book

	//encode json
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)

}
