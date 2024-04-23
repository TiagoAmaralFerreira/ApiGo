package handlers

import (
	"crudPostgreSQL/models"
	"encoding/json"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)

}
