package main

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func filterByStatus(notes []note, status bool) []note {
	notesCopy := []note{}
	for _, note := range notes {
		if note.IsDone == status {
			notesCopy = append(notesCopy, note)
		}
	}

	return notesCopy
}

func filterByText(notes []note, text string) []note {
	notesCopy := []note{}
	for _, note := range notes {
		if strings.Contains(note.Text, text) {
			notesCopy = append(notesCopy, note)
		}
	}

	return notesCopy
}

func filterByTime(notes []note, minimum time.Time, maximum time.Time) []note {
	notesCopy := []note{}
	for _, note := range notes {
		tooEarly := note.CreatedAt.Before(minimum)
		tooLate := note.CreatedAt.After(maximum)
		if !tooEarly && !tooLate {
			notesCopy = append(notesCopy, note)
		}
	}

	return notesCopy
}

func filterByDate(notes []note, date time.Time) []note {
	minimum := date.Truncate(24 * time.Hour)
	maximum := minimum.Add(24 * time.Hour)
	return filterByTime(notes, minimum, maximum)
}

func filterByCommand(notes []note, line string) ([]note, error) {
	filteredNotes := []note{}
	if strings.HasPrefix(line, "list") {
		parameter := getParameter(line, "list")
		if parameter == "done" {
			filteredNotes = filterByStatus(notes, true)
		} else if parameter == "to do" {
			filteredNotes = filterByStatus(notes, false)
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

		filteredNotes = filterByText(notes, query)
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

		filteredNotes = filterByDate(notes, date)
	}

	return filteredNotes, nil
}

func filterByMultiCommand(notes []note, line string) ([]note, error) {
	commands := strings.Split(line, "|")
	for index, command := range commands {
		filteredNotes, err := filterByCommand(notes, command)
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
