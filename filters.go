package main

func filterByStatus(notes []note, status bool) []note {
	notesCopy := []note{}
	for _, note := range notes {
		if note.IsDone == status {
			notesCopy = append(notesCopy, note)
		}
	}

	return notesCopy
}
