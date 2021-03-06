package main

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func unmarshalNote(line string) (note, error) {
	parts := strings.SplitN(line, " ", 8)
	if len(parts) < 7 {
		return note{}, errors.New("not enough parts in the line")
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return note{}, errors.New("unable to parse the note ID: " + err.Error())
	}
	id = id - idOffsetForAlignment

	createdAtParts := []string{parts[2], parts[3], parts[4], parts[5], parts[6]}
	createdAtStr := strings.Join(createdAtParts, " ")
	createdAt, err := time.Parse(time.RFC822Z, createdAtStr)
	if err != nil {
		return note{},
			errors.New("unable to parse the note creation timestamp: " + err.Error())
	}

	isDone := parts[1] == "[x]"

	note := note{ID: id, CreatedAt: createdAt, IsDone: isDone, Text: parts[7]}
	return note, nil
}

func unmarshalNotes(text string) ([]note, error) {
	notes := []note{}
	lines := strings.Split(text, "\n")
	for lineIndex, line := range lines {
		if line == "" {
			continue
		}

		note, err := unmarshalNote(line)
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
