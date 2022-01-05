package handellers

import (
	"encoding/json"
	"fmt"
	"github.com/iamtakdir/crud-go/pkg/models"
	"net/http"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {

	var book []models.Book

	if result := h.DB.Find(&book); result.Error != nil {
		fmt.Println(result.Error)
	}
	
	fmt.Println(book)
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)

}
