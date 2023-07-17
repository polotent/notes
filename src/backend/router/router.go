package router

import (
	"backend/controller"
	"backend/db"
	"backend/repository"
	"backend/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Init(mux *http.ServeMux) {
	setNoteController(mux)
}

func setNoteController(mux *http.ServeMux) {
	nr := repository.NewNoteRepository(db.DB)
	ns := service.NewNoteService(nr)
	nc := controller.NewNoteController(ns)

	mux.HandleFunc(fmt.Sprintf("%s/notes", API_PREFIX), func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			sendJsonResponse(w, http.StatusMethodNotAllowed, []byte(`{
				"success": false,
				"message": "Invalid HTTP method for requested endpoint",
			}`))
			log.Printf("Invalid HTTP method for requested endpoint: %s", r.URL)
			return
		}

		if r.Method == http.MethodGet {
			notes, err := nc.ReadAllNotes()
			if err != nil {
				sendJsonResponse(w, http.StatusInternalServerError, []byte(`{
					"success": false,
					"message": "Couldn't read from database",
				}`))
				log.Printf("Note controller error: %s", err)
				return
			}
		
			jsonResp, err := json.Marshal(notes)

			if err != nil {
				sendJsonResponse(w, http.StatusInternalServerError, []byte(`{
					"success": false,
					"message": "Error happened while marshalling response",
				}`))
				log.Printf("JSON marshalling error: %s", err)
				return
			}

			sendJsonResponse(w, http.StatusOK, []byte(fmt.Sprintf(`{
				"success": true,
				"data": %s
			}`, jsonResp)))
		}
	})
}
