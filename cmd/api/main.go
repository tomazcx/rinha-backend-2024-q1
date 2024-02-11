package main

import (
	"log"
	"net/http"

	"github.com/tomazcx/rinha-backend-2024-q1/config"
	"github.com/tomazcx/rinha-backend-2024-q1/internal/httpapp"
)

func main(){
	mux := http.NewServeMux()

	router := httpapp.Router{}
	router.Init(mux)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading the config: %v", err)
	}

	db, err := config.ConnectToDb(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	log.Println("Server started at port " + cfg.WebPort)
	http.ListenAndServe(":8000", mux)
}
