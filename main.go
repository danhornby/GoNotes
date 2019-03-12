package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//The Note Type Object
type Note struct {
	ID      string `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

var notes []Note

//GetNotes function
func GetNotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(notes)
}

//GetNote function
func GetNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range notes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Note{})
}

//CreateNote function
func CreateNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)
	note.ID = params["id"]
	notes = append(notes, note)
	json.NewEncoder(w).Encode(notes)
}

//DeleteNote Function
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range notes {
		if item.ID == params["id"] {
			notes = append(notes[:index], notes[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(notes)
	}
}

//Main Execution Function
func main() {
	router := mux.NewRouter()
	notes = append(notes, Note{ID: "1", Title: "Does this record?", Content: "No one will see this."})
	router.HandleFunc("/Notes", GetNotes).Methods("GET")
	router.HandleFunc("/Notes/{id}", GetNote).Methods("GET")
	router.HandleFunc("/Notes/{id}", CreateNote).Methods("POST")
	router.HandleFunc("/Notes/{id})", DeleteNote).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8054", router))
}

