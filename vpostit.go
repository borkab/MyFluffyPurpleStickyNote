package vpostit

// Note represents an online post-it
type Note struct {
	Title string //title of your note
	Body  string //your sticky note
	Info
	ID string //a unique identifier for your note
}

type Info struct {
	MadeDay    string
	LastChange string
}
