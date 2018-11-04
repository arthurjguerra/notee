package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var content string
	_ = json.NewDecoder(r.Body).Decode(&content)
	note := GetNote(id, content)

	json.NewEncoder(w).Encode(note)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{id}", GetNoteHandler).Methods("GET", "POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
