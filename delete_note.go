package todo

// DeleteNote ...
func DeleteNote(notes []Note, id int) []Note {
	notesCopy := []Note{}
	for _, note := range notes {
		if note.ID != id {
			notesCopy = append(notesCopy, note)
		}
	}

	return notesCopy
}
