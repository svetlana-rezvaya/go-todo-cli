package main

import (
	"errors"
	"strings"
	"time"
)

func updateStatus(notes []note, id int, status bool) {
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

	notes[targetIndex].UpdatedAt = time.Now()
	notes[targetIndex].IsDone = status
}

func updateUsingCommand(notes []note, line string) ([]note, error) {
	if strings.HasPrefix(line, "add") {
		text := getParameter(line, "add")
		if text == "" {
			return nil, errors.New("text missing in 'add' command")
		}

		notes = createNote(notes, text)
	} else if strings.HasPrefix(line, "check") {
		id, err := getIDParameter(line, "check")
		if err != nil {
			return nil, errors.New("unable to get the note ID: " + err.Error())
		}

		updateStatus(notes, id, true)
	} else if strings.HasPrefix(line, "uncheck") {
		id, err := getIDParameter(line, "uncheck")
		if err != nil {
			return nil, errors.New("unable to get the note ID: " + err.Error())
		}

		updateStatus(notes, id, false)
	} else if strings.HasPrefix(line, "delete") {
		id, err := getIDParameter(line, "delete")
		if err != nil {
			return nil, errors.New("unable to get the note ID: " + err.Error())
		}

		notes = deleteNote(notes, id)
	}

	return notes, nil
}
