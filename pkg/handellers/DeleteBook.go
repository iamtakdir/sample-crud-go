package handellers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/iamtakdir/crud-go/pkg/models"
	"net/http"
	"strconv"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	//get the unique id from url
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	var book models.Book
	// find a book by id
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	//Delete that book

	h.DB.Delete(&book)

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
