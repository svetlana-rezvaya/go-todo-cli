package main

import "strings"

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
