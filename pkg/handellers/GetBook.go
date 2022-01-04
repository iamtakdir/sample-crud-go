package handellers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/iamtakdir/crud-go/pkg/mock"
	"net/http"
	"strconv"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	//Read dynamic id parameter
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	//	itarate over books
	for _, book := range mock.Books {
		if book.Id == id {
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
		}
	}

}
