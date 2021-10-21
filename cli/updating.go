package cli

import (
	"errors"
	"strings"

	todo "github.com/svetlana-rezvaya/go-todo-cli"
)

// UpdateUsingCommand ...
func UpdateUsingCommand(notes []todo.Note, line string) ([]todo.Note, error) {
	if strings.HasPrefix(line, "add") {
		text := getParameter(line, "add")
		if text == "" {
			return nil, errors.New("text missing in 'add' command")
		}

		notes = todo.CreateNote(notes, text)
	} else if strings.HasPrefix(line, "check") {
		id, err := getIDParameter(line, "check")
		if err != nil {
			return nil, errors.New("unable to get the note ID: " + err.Error())
		}

		todo.UpdateStatus(notes, id, true)
	} else if strings.HasPrefix(line, "uncheck") {
		id, err := getIDParameter(line, "uncheck")
		if err != nil {
			return nil, errors.New("unable to get the note ID: " + err.Error())
		}

		todo.UpdateStatus(notes, id, false)
	} else if strings.HasPrefix(line, "delete") {
		id, err := getIDParameter(line, "delete")
		if err != nil {
			return nil, errors.New("unable to get the note ID: " + err.Error())
		}

		notes = todo.DeleteNote(notes, id)
	}

	return notes, nil
}
