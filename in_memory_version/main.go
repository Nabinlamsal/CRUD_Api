package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nabinlamsal/crudApi/models"
	"github.com/nabinlamsal/crudApi/router"
)

// fake db
func main() {
	fmt.Println("Building APIs in Golag!")

	router.SetUpRouter()

	models.SeedNotes()

	//listen to the server or port
	log.Fatal(http.ListenAndServe(":5000", router.SetUpRouter()))
}
