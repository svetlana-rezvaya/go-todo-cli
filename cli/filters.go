package cli

import (
	"errors"
	"fmt"
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
			return nil, fmt.Errorf("unknown parameter for 'list' command: %s", parameter)
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
				fmt.Errorf("unable to parse the 'date' command parameter: %s", err)
		}

		filteredNotes = todo.FilterByDate(notes, date)
	} else {
		return nil, fmt.Errorf("unknown command: %s", line)
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
			return nil, fmt.Errorf("unable to filter by command #%d: %s", index+1, err)
		}

		notes = filteredNotes
	}

	return notes, nil
}
