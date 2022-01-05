package handellers

import (
	"encoding/json"
	"github.com/iamtakdir/crud-go/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func GetBook(w http.ResponseWriter, r *http.Request) {

	dsn := "root:root@tcp(127.0.0.1:3306)/bookstoredb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//Read dynamic id parameter
	//vars := mux.Vars(r)
	//
	//id, _ := strconv.Atoi(vars["id"])

	var book models.Book
	// Get all records
	result := db.Find(&book)
	// SELECT * FROM users;

	//encode json
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
