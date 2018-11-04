package main

import (
	"errors"
)

type Note struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func listNotes() []Note {
	var notes []Note
	return notes
}

func findNote(id string) (Note, error) {
	notes := listNotes()

	for _, n := range notes {
		if n.ID == id {
			return n, nil
		}
	}

	return Note{}, errors.New("Note not found")
}

func GetNote(id string, content string) Note {
	note, err := findNote(id)

	if err != nil {
		note = CreateNote(id, "")
	}

	return note
}

// create a new item
func CreateNote(id string, content string) Note {
	var note Note
	note.ID = id
	note.Content = content
	return note
}
