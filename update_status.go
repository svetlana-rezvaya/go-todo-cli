package main

func updateStatus(notes []note, id int, status bool) {
	targetIndex := -1
	for index, note := range notes {
		if note.ID == id {
			targetIndex = index
			break
		}
	}
	if targetIndex == -1 {
		return
	}

	notes[targetIndex].IsDone = status
}
