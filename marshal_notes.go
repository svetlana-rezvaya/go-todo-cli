package main

import "strconv"

func marshalNote(note note) string {
	// add an offset for all ids to have the same width
	id := strconv.Itoa(note.ID + 10000)

	flag := ""
	if note.IsDone {
		flag = "[x]"
	} else {
		flag = "[ ]"
	}

	return id + " " + flag + " " + note.Text
}

func marshalNotes(notes []note) string {
	result := ""
	for _, note := range notes {
		line := marshalNote(note)
		result = result + line + "\n"
	}

	return result
}
