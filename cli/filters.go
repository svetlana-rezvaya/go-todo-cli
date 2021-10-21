package cli

import (
	"errors"
	"strconv"
	"strings"
	"time"

	todo "github.com/svetlana-rezvaya/go-todo-cli"
)

// FilterByCommand ...
func FilterByCommand(notes []todo.Note, line string) ([]todo.Note, error) {
	filteredNotes := []todo.Note{}
	if strings.HasPrefix(line, "list") {
		parameter := getParameter(line, "list")
		if parameter == "done" {
			filteredNotes = todo.FilterByStatus(notes, true)
		} else if parameter == "to do" {
			filteredNotes = todo.FilterByStatus(notes, false)
		} else if parameter == "" {
			filteredNotes = notes
		} else {
			return nil, errors.New("unknown parameter for 'list' command: " + parameter)
		}
	} else if strings.HasPrefix(line, "find") {
		query := getParameter(line, "find")
		if query == "" {
			return nil, errors.New("query missing in 'find' command")
		}

		filteredNotes = todo.FilterByText(notes, query)
	} else if strings.HasPrefix(line, "date") {
		parameter := getParameter(line, "date")
		if parameter == "" {
			return nil, errors.New("parameter missing in 'date' command")
		}

		date, err := time.Parse("02 Jan 06", parameter)
		if err != nil {
			return nil,
				errors.New("unable to parse the 'date' command parameter: " + err.Error())
		}

		filteredNotes = todo.FilterByDate(notes, date)
	} else {
		return nil, errors.New("unknown command: " + line)
	}

	return filteredNotes, nil
}

// FilterByMultiCommand ...
func FilterByMultiCommand(notes []todo.Note, line string) ([]todo.Note, error) {
	commands := strings.Split(line, "|")
	for index, command := range commands {
		command = strings.TrimSpace(command)
		if command == "" {
			continue
		}

		filteredNotes, err := FilterByCommand(notes, command)
		if err != nil {
			return nil, errors.New(
				"unable to filter by command #" + strconv.Itoa(index+1) + ": " +
					err.Error(),
			)
		}

		notes = filteredNotes
	}

	return notes, nil
}
