package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Nabinlamsal/fiber-postgress/models"
	"github.com/Nabinlamsal/fiber-postgress/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

//	type Notes struct {
//		Title   string `json:"title"`
//		Content string `json:"content"`
//		Creator string `json:"creator"`
//	}
type Repo struct {
	DB *gorm.DB
}

func (r *Repo) CreateNotes(context *fiber.Ctx) error {
	note := models.Notes{}

	//if http is used, w and r are used....so they are running in background here
	err := context.BodyParser(&note)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&note).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not Create Notes"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "notes created successfully"})
	return nil //everything is going well and the prohram will return nil

}

func (r *Repo) GetAllNotes(context *fiber.Ctx) error {
	notesModels := []models.Notes{}

	err := r.DB.Find(&notesModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not retrive books!"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "notes retrieved successfully!",
			"data":    notesModels,
		})
	return nil
}

func (r *Repo) DeleteNotes(context *fiber.Ctx) error {
	notesModel := models.Notes{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "id cant be empty",
			})
		return nil
	}
	err := r.DB.Delete(&notesModel, id)
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete Notes",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "notes deleted successfully",
	})
	return nil

}

func (r *Repo) GetNotes(context *fiber.Ctx) error {
	notesModel := models.Notes{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "id cant be empty",
			})
		return nil
	}
	err := r.DB.Where("id = ?", id).First(&notesModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get the notes from id",
		})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "notes retrieved successfully",
		"data":    notesModel,
	})
	return nil
}

func (r *Repo) SetupRoutes(app *fiber.App) {
	api := app.Group("api")
	api.Post("/create_notes", r.CreateNotes)
	api.Delete("/delete_notes/:id", r.DeleteNotes)
	api.Get("/get_notes/:id", r.GetNotes)
	api.Get("/notes", r.GetAllNotes)

}

func main() {
	fmt.Println("Eposide on Go and Postgress Connection!")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(&config)
	if err != nil {
		log.Fatal("Could not connect to the Database!")
	}

	if err := models.MigrateNotes(db); err != nil {
		log.Fatal("Could not migrate database!")
	}

	r := Repo{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8000")
}

// The Repo struct holds a *gorm.DB object, which is the actual connection to PostgreSQL.
// The r variable is an instance of Repo.
// When we call r.CreateNotes(), it's a method that internally uses r.DB,
// which is the GORM engine doing the real database work.
// This pattern creates a clean abstraction: you call methods on Repo, and GORM handles the database behind the scenes.
