package flashcards

type FlashCard struct {
	Question string
	Answer string
	Subject string
	Hint string
	Level int
	Tags []string
}

// write to json file.

// load from json file.