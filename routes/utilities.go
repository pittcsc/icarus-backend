// Package routes provides ...
package routes

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
)

var db gorm.DB

func InitializeDBConnection(dbconnection gorm.DB) {
	db = dbconnection
}

func sendJSON(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	json.NewEncoder(w).Encode(obj)
}

func sendUnprocessableEntity(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422) // unprocessable entity
	if err := json.NewEncoder(w).Encode(err); err != nil {
		panic(err)
	}
}
