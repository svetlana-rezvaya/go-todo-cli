package main

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// FilterByStatus ...
func FilterByStatus(notes []Note, status bool) []Note {
	notesCopy := []Note{}
	for _, note := range notes {
		if note.IsDone == status {
			notesCopy = append(notesCopy, note)
		}
	}

	return notesCopy
}

// FilterByText ...
func FilterByText(notes []Note, text string) []Note {
	notesCopy := []Note{}
	for _, note := range notes {
		if strings.Contains(note.Text, text) {
			notesCopy = append(notesCopy, note)
		}
	}

	return notesCopy
}

// FilterByTime ...
func FilterByTime(notes []Note, minimum time.Time, maximum time.Time) []Note {
	notesCopy := []Note{}
	for _, note := range notes {
		tooEarly := note.CreatedAt.Before(minimum)
		tooLate := note.CreatedAt.After(maximum)
		if !tooEarly && !tooLate {
			notesCopy = append(notesCopy, note)
		}
	}

	return notesCopy
}

// FilterByDate ...
func FilterByDate(notes []Note, date time.Time) []Note {
	minimum := date.Truncate(24 * time.Hour)
	maximum := minimum.Add(24 * time.Hour)
	return FilterByTime(notes, minimum, maximum)
}

// FilterByCommand ...
func FilterByCommand(notes []Note, line string) ([]Note, error) {
	filteredNotes := []Note{}
	if strings.HasPrefix(line, "list") {
		parameter := getParameter(line, "list")
		if parameter == "done" {
			filteredNotes = FilterByStatus(notes, true)
		} else if parameter == "to do" {
			filteredNotes = FilterByStatus(notes, false)
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

		filteredNotes = FilterByText(notes, query)
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

		filteredNotes = FilterByDate(notes, date)
	} else {
		return nil, errors.New("unknown command: " + line)
	}

	return filteredNotes, nil
}

// FilterByMultiCommand ...
func FilterByMultiCommand(notes []Note, line string) ([]Note, error) {
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
