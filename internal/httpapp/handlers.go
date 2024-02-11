package httpapp

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/tomazcx/rinha-backend-2024-q1/internal/db"
)

type Router struct{}

func (r *Router) Init(m *http.ServeMux) {
	m.HandleFunc("POST /clientes/{id}/transacoes", HandleCreateTransaction)

	m.HandleFunc("GET /clientes/{id}/extrato", HandleGetExtract)
}

func HandleCreateTransaction(w http.ResponseWriter, r *http.Request) {
	var body db.CreateTransactionDTO 

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if body.Tipo != "c" && body.Tipo != "d" {		
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pathId := r.PathValue("id")
	clientId, err := strconv.Atoi(pathId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if clientId < 1 || clientId > 5 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	repo := db.NewClientRepository()
	clientData, err := repo.CreateTransaction(clientId, body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error creating transaction: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clientData)
}

func HandleGetExtract(w http.ResponseWriter, r *http.Request){
	pathId := r.PathValue("id")
	clientId, err := strconv.Atoi(pathId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if clientId < 1 || clientId > 5 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	repo := db.NewClientRepository()
	extract, err := repo.GetExtract(clientId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error creating transaction: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(extract)
}
