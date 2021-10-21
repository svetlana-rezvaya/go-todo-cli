package todo

import (
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
