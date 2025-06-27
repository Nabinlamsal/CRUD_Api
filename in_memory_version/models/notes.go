package models

// model for courses - file
type Notes struct {
	NotesId      string   `json:"notesid"`
	NotesTitle   string   `json:"title"`
	NotesContent string   `json:"content"`
	Creator      *Creator `json:"creator"`
}

type Creator struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake db - Made it Exported (capitalized) so other packages like controllers can access it
var NotesList []Notes

// middleware, helper -file
func (n *Notes) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return n.NotesTitle == "" // i don't want to check for the course ID, I want that user should be allowed to move forward if the course id is not empty, I don't want to rely on the user so wanna create manually
}

func SeedNotes() {
	// seeding
	NotesList = append(NotesList,
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
}
