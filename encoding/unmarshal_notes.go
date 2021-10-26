package encoding

import (
	"errors"
	"strconv"
	"strings"
	"time"

	todo "github.com/svetlana-rezvaya/go-todo-cli"
)

// UnmarshalNote ...
func UnmarshalNote(line string) (todo.Note, error) {
	parts := strings.SplitN(line, " ", 8)
	if len(parts) < 7 {
		return todo.Note{}, errors.New("not enough parts in the line")
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return todo.Note{}, errors.New("unable to parse the note ID: " + err.Error())
	}
	id = id - IDOffsetForAlignment

	createdAtParts := []string{parts[2], parts[3], parts[4], parts[5], parts[6]}
	createdAtStr := strings.Join(createdAtParts, " ")
	createdAt, err := time.Parse(time.RFC822Z, createdAtStr)
	if err != nil {
		return todo.Note{},
			errors.New("unable to parse the note creation timestamp: " + err.Error())
	}

	isDone := parts[1] == "[x]"

	note := todo.Note{ID: id, CreatedAt: createdAt, IsDone: isDone, Text: parts[7]}
	return note, nil
}

// UnmarshalNotes ...
func UnmarshalNotes(text string) ([]todo.Note, error) {
	notes := []todo.Note{}
	lines := strings.Split(text, "\n")
	for lineIndex, line := range lines {
		if line == "" {
			continue
		}

		note, err := UnmarshalNote(line)
		if err != nil {
			lineIndexStr := strconv.Itoa(lineIndex + 1)
			return nil, errors.New(
				"unable to unmarshal the note in line #" + lineIndexStr + ": " +
					err.Error(),
			)
		}

		notes = append(notes, note)
	}

	return notes, nil
}
