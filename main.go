package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for courses - file
type Notes struct {
	NotesId      string `json:"notesid"`
	NotesTitle   string `json:"title"`
	NotesContent string `json:"content"`
	Creator      *Creator
}

type Creator struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var notes []Notes

// middleware, helper -file
func (n *Notes) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return n.NotesTitle == "" //i dont want to check for the course ID, i want that user should be allowed to move forward if the corse id is not empty, i dont want to rely on the user so wanna create manually
}
func main() {
	fmt.Println("Building APIs in Golag!")
	r := mux.NewRouter()

	//seeding
	notes = append(notes,
		Notes{
			NotesId:      "2",
			NotesTitle:   "Golang",
			NotesContent: "This is Notes related to GO",
			Creator:      &Creator{Fullname: "Nabin Lamsal", Website: "www.example.com"},
		},
		Notes{
			NotesId:      "4",
			NotesTitle:   "ReactJS",
			NotesContent: "This is Notes related to ReactJS",
			Creator:      &Creator{Fullname: "Rajesh Hamal", Website: "www.trishul.com"},
		},
	)

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/notes", getAllNotes).Methods("GET")
	r.HandleFunc("/note/{id}", getOneNote).Methods("GET") //params is governed
	r.HandleFunc("/note", createOneNote).Methods("POST")
	r.HandleFunc("/note/{id}", updateOneNote).Methods("PUT")
	r.HandleFunc("/note/{id}", deleteOneNote).Methods("DELETE")
	//listen to the server or port
	log.Fatal(http.ListenAndServe(":5000", r))
}

//controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API and this is Home Section</>"))
}

func getAllNotes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Notes")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func getOneNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//looping through the courses, finding matching IDs and returning the responces
	for _, note := range notes {
		if note.NotesId == params["id"] {
			json.NewEncoder(w).Encode(note)
			return
		}
	}
	json.NewEncoder(w).Encode("No Id Found")
	return
}

func createOneNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Notes")
	w.Header().Set("Content-Type", "application/json")

	//what if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data!")
	}

	//what about {}
	var note Notes
	_ = json.NewDecoder(r.Body).Decode(&note)
	if note.IsEmpty() {
		_ = json.NewEncoder(w).Encode("No data inside JSON!")
		return
	}

	//generating random ids, string
	//append new course in courses
	rand.Seed(time.Now().UnixNano())
	note.NotesId = strconv.Itoa(rand.Intn(100))
	notes = append(notes, note)
	json.NewEncoder(w).Encode(note)
	return

}

func updateOneNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Notes")
	w.Header().Set("Content-Type", "application/json")

	//grab id
	params := mux.Vars(r)

	//loop through the index, remove, add with my id
	for index, note := range notes {

		if note.NotesId == params["id"] {
			var note Notes
			//skipping note at index
			notes = append(notes[:index], notes[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&note)
			//  Since we removed the note, and IDs should be unique, we reassign the same ID to the new data.
			note.NotesId = params["id"]
			// Add the updated note back to the notes list.
			notes = append(notes, note)
			// Send the updated note back to the client as the response.
			json.NewEncoder(w).Encode(note)
			return
		}
	}
	json.NewEncoder(w).Encode("Id not found")
}
func deleteOneNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Notes")
	w.Header().Set("Content-Type", "application/json")

	//grab the id
	params := mux.Vars(r)
	deleted := false
	//loop, id, remove, index, index+1
	for index, note := range notes {
		if note.NotesId == params["id"] {
			notes = append(notes[:index], notes[index+1:]...) //removed
			deleted = true
			break
		}
	}
	if deleted {
		json.NewEncoder(w).Encode("Notes Deleted Succcessfully")
	}
}
