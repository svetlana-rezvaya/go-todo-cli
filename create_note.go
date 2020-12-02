package main

import "strings"

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
	text = strings.TrimSpace(text)
	note := note{ID: id, IsDone: false, Text: text}
	return append(notes, note)
}
