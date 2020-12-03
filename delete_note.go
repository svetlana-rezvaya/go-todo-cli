package main

func deleteNote(notes []note, id int) []note {
	notesCopy := []note{}
	for _, note := range notes {
		if note.ID != id {
			notesCopy = append(notesCopy, note)
		}
	}

	return notesCopy
}
