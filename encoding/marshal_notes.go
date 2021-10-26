package encoding

import (
	"strconv"
	"time"

	todo "github.com/svetlana-rezvaya/go-todo-cli"
)

// IDOffsetForAlignment ...
const IDOffsetForAlignment = 10000

// MarshalNote ...
func MarshalNote(note todo.Note) string {
	// add an offset for all ids to have the same width
	id := strconv.Itoa(note.ID + IDOffsetForAlignment)

	flag := ""
	if note.IsDone {
		flag = "[x]"
	} else {
		flag = "[_]"
	}

	createdAt := note.CreatedAt.Format(time.RFC822Z)

	return id + " " + flag + " " + createdAt + " " + note.Text
}

// MarshalNotes ...
func MarshalNotes(notes []todo.Note) string {
	result := ""
	for _, note := range notes {
		line := MarshalNote(note)
		result = result + line + "\n"
	}

	return result
}
