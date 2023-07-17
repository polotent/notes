package main

import (
	"log"
	"net/http"
	"backend/router"
	"backend/db"
)

func main() {
	defer func() {
		if err := db.DB.Close(); err != nil {
			log.Fatalf("Database connection close failed: %s", err)
		}
	}()

	mux := http.NewServeMux()
	router.Init(mux)
	server := &http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server launch failed: %s", err)
	}
}
