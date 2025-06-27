package router

import (
	"github.com/gorilla/mux"
	"github.com/nabinlamsal/crudApi/controllers"
)

func SetUpRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.ServeHome).Methods("GET")
	r.HandleFunc("/notes", controllers.GetAllNotes).Methods("GET")
	r.HandleFunc("/note/{id}", controllers.GetOneNote).Methods("GET") //params is governed
	r.HandleFunc("/note", controllers.CreateOneNote).Methods("POST")
	r.HandleFunc("/note/{id}", controllers.UpdateOneNote).Methods("PUT")
	r.HandleFunc("/note/{id}", controllers.DeleteOneNote).Methods("DELETE")
	return r

}
