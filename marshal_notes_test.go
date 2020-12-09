package main

import (
	"testing"
	"time"
)

func Test_marshalNote_isNotDone(test *testing.T) {
	createdAt := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	updatedAt := createdAt.Add(567 * time.Hour)
	note := note{
		ID:        123,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		IsDone:    false,
		Text:      "test",
	}
	text := marshalNote(note)

	if text != "10123 [ ] 09 Dec 20 21:05 +0000 test" {
		test.Fail()
	}
}

func Test_marshalNote_isDone(test *testing.T) {
	createdAt := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	updatedAt := createdAt.Add(567 * time.Hour)
	note := note{
		ID:        123,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		IsDone:    true,
		Text:      "test",
	}
	text := marshalNote(note)

	if text != "10123 [x] 09 Dec 20 21:05 +0000 test" {
		test.Fail()
	}
}

func Test_marshalNotes_empty(test *testing.T) {
	text := marshalNotes([]note{})

	if text != "" {
		test.Fail()
	}
}

func Test_marshalNotes_nonEmpty(test *testing.T) {
	createdAt := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	notes := []note{
		note{ID: 100, CreatedAt: createdAt, IsDone: false, Text: "one"},
		note{ID: 101, CreatedAt: createdAt.Add(time.Hour), IsDone: true, Text: "two"},
		note{
			ID:        102,
			CreatedAt: createdAt.Add(2 * time.Hour),
			IsDone:    false,
			Text:      "three",
		},
		note{
			ID:        103,
			CreatedAt: createdAt.Add(3 * time.Hour),
			IsDone:    true,
			Text:      "four",
		},
		note{
			ID:        104,
			CreatedAt: createdAt.Add(4 * time.Hour),
			IsDone:    false,
			Text:      "five",
		},
	}
	text := marshalNotes(notes)

	wantedText := "10100 [ ] 09 Dec 20 21:05 +0000 one\n" +
		"10101 [x] 09 Dec 20 22:05 +0000 two\n" +
		"10102 [ ] 09 Dec 20 23:05 +0000 three\n" +
		"10103 [x] 10 Dec 20 00:05 +0000 four\n" +
		"10104 [ ] 10 Dec 20 01:05 +0000 five\n"
	if text != wantedText {
		test.Fail()
	}
}
