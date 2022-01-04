package handellers

import (
	"encoding/json"
	"github.com/iamtakdir/crud-go/pkg/mock"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mock.Books)
}
