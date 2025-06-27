package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/nabinlamsal/crudApi/models"
)

// controllers - file
// Removed local notes slice to avoid confusion, using models.NotesList instead

// serve home route
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API and this is Home Section</>"))
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Notes")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.NotesList) // use models.NotesList here
}

func GetOneNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//looping through the courses, finding matching IDs and returning the responces
	for _, note := range models.NotesList {
		if note.NotesId == params["id"] {
			json.NewEncoder(w).Encode(note) // encode note, not notes
			return
		}
	}
	json.NewEncoder(w).Encode("No Id Found")
	// return
}

func CreateOneNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Notes")
	w.Header().Set("Content-Type", "application/json")

	//what if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data!")
		return // add return to avoid continuing
	}

	//what about {}
	var note models.Notes
	_ = json.NewDecoder(r.Body).Decode(&note)
	if note.IsEmpty() {
		_ = json.NewEncoder(w).Encode("No data inside JSON!")
		return
	}

	//generating random ids, string
	//append new course in courses
	rand.Seed(time.Now().UnixNano())
	note.NotesId = strconv.Itoa(rand.Intn(100))
	models.NotesList = append(models.NotesList, note) // use models.NotesList here
	json.NewEncoder(w).Encode(note)
	return
}

func UpdateOneNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Notes")
	w.Header().Set("Content-Type", "application/json")

	//grab id
	params := mux.Vars(r)

	//loop through the index, remove, add with my id
	for index, note := range models.NotesList { // use models.NotesList here

		if note.NotesId == params["id"] {
			var note models.Notes
			//skipping note at index
			models.NotesList = append(models.NotesList[:index], models.NotesList[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&note)
			//  Since we removed the note, and IDs should be unique, we reassign the same ID to the new data.
			note.NotesId = params["id"]
			// Add the updated note back to the notes list.
			models.NotesList = append(models.NotesList, note)
			// Send the updated note back to the client as the response.
			json.NewEncoder(w).Encode(note)
			return
		}
	}
	json.NewEncoder(w).Encode("Id not found")
}

func DeleteOneNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Notes")
	w.Header().Set("Content-Type", "application/json")

	//grab the id
	params := mux.Vars(r)
	deleted := false
	//loop, id, remove, index, index+1
	for index, note := range models.NotesList { // use models.NotesList here
		if note.NotesId == params["id"] {
			models.NotesList = append(models.NotesList[:index], models.NotesList[index+1:]...) //removed
			deleted = true
			break
		}
	}
	if deleted {
		json.NewEncoder(w).Encode("Notes Deleted Succcessfully")
	}
}
