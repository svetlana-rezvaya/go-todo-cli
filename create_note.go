package main

func getMaximalID(notes []note) int {
	maximalID := 0
	for _, note := range notes {
		if note.ID > maximalID {
			maximalID = note.ID
		}
	}

	return maximalID
}
