package main

import (
	"strings"
	"time"
)

func getMaximalID(notes []note) int {
	maximalID := 0
	for _, note := range notes {
		if note.ID > maximalID {
			maximalID = note.ID
		}
	}

	return maximalID
}

func createNote(notes []note, text string) []note {
	id := getMaximalID(notes) + 1
	currentTime := time.Now()
	text = strings.TrimSpace(text)
	note := note{
		ID:        id,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		IsDone:    false,
		Text:      text,
	}
	return append(notes, note)
}
