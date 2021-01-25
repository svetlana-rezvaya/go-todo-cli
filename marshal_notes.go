package main

import (
	"strconv"
	"time"
)

const idOffsetForAlignment = 10000

func marshalNote(note note) string {
	// add an offset for all ids to have the same width
	id := strconv.Itoa(note.ID + idOffsetForAlignment)

	flag := ""
	if note.IsDone {
		flag = "[x]"
	} else {
		flag = "[_]"
	}

	createdAt := note.CreatedAt.Format(time.RFC822Z)

	return id + " " + flag + " " + createdAt + " " + note.Text
}

func marshalNotes(notes []note) string {
	result := ""
	for _, note := range notes {
		line := marshalNote(note)
		result = result + line + "\n"
	}

	return result
}
