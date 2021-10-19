package main

import (
	"strings"
	"time"
)

// CreateNote ...
func CreateNote(notes []Note, text string) []Note {
	id := getMaximalID(notes) + 1
	currentTime := time.Now()
	text = strings.TrimSpace(text)
	note := Note{
		ID:        id,
		CreatedAt: currentTime,
		IsDone:    false,
		Text:      text,
	}
	return append(notes, note)
}

func getMaximalID(notes []Note) int {
	maximalID := 0
	for _, note := range notes {
		if note.ID > maximalID {
			maximalID = note.ID
		}
	}

	return maximalID
}
