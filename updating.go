package todo

import (
	"errors"
	"strings"
)

// UpdateStatus ...
func UpdateStatus(notes []Note, id int, status bool) {
	targetIndex := -1
	for index, note := range notes {
		if note.ID == id {
			targetIndex = index
			break
		}
	}
	if targetIndex == -1 {
		return
	}

	notes[targetIndex].IsDone = status
}

// UpdateUsingCommand ...
func UpdateUsingCommand(notes []Note, line string) ([]Note, error) {
	if strings.HasPrefix(line, "add") {
		text := getParameter(line, "add")
		if text == "" {
			return nil, errors.New("text missing in 'add' command")
		}

		notes = CreateNote(notes, text)
	} else if strings.HasPrefix(line, "check") {
		id, err := getIDParameter(line, "check")
		if err != nil {
			return nil, errors.New("unable to get the note ID: " + err.Error())
		}

		UpdateStatus(notes, id, true)
	} else if strings.HasPrefix(line, "uncheck") {
		id, err := getIDParameter(line, "uncheck")
		if err != nil {
			return nil, errors.New("unable to get the note ID: " + err.Error())
		}

		UpdateStatus(notes, id, false)
	} else if strings.HasPrefix(line, "delete") {
		id, err := getIDParameter(line, "delete")
		if err != nil {
			return nil, errors.New("unable to get the note ID: " + err.Error())
		}

		notes = DeleteNote(notes, id)
	}

	return notes, nil
}
