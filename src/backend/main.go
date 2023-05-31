package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type Note struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var notes []Note = []Note{
	{1, "First", "Beautiful"},
	{2, "Second", "Crazy"},
}

func main() {
	err := DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	var res Note
	var notes []Note
	rows, err := DB.Query("SELECT * FROM notes")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		if err := rows.Scan(&res.Id, &res.Title, &res.Description); err != nil {
			log.Fatal(err)
		}
		notes = append(notes, res)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(notes)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now(), " - ", r.Method, " - ", r.URL)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(notes)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}
