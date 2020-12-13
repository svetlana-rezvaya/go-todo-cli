package main

import (
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
