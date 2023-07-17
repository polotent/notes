package router

import (
	"log"
	"net/http"
)

func sendJsonResponse(w http.ResponseWriter, statusCode int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(data)
	if err != nil {
		log.Printf("Error while sending http response: %s", err)
	}
}