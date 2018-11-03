package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Note struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

var notes []Note

func findNote(id string) (Note, error) {
	for _, n := range notes {
		if n.ID == id {
			return n, nil
		}
	}

	return Note{}, errors.New("Note not found")
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	note, err := findNote(params["id"])
	if err != nil {
		note = CreateNote(w, r)
	}

	json.NewEncoder(w).Encode(note)
}

// create a new item
func CreateNote(w http.ResponseWriter, r *http.Request) Note {
	params := mux.Vars(r)
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)
	note.ID = params["id"]
	notes = append(notes, note)
	return note
}

func main() {
	notes = append(notes, Note{ID: "mynote", Content: "this is my first notepad online"})

	r := mux.NewRouter()

	r.HandleFunc("/{id}", GetNote).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
