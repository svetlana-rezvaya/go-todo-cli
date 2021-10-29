package encoding

import (
	"testing"
	"time"

	todo "github.com/svetlana-rezvaya/go-todo-cli"
)

func TestMarshalNote_isNotDone(test *testing.T) {
	createdAt := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	note := todo.Note{
		ID:        123,
		CreatedAt: createdAt,
		IsDone:    false,
		Text:      "test",
	}
	text := MarshalNote(note)

	if text != "10123 [_] 09 Dec 20 21:05 +0000 test" {
		test.Fail()
	}
}

func TestMarshalNote_isDone(test *testing.T) {
	createdAt := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	note := todo.Note{
		ID:        123,
		CreatedAt: createdAt,
		IsDone:    true,
		Text:      "test",
	}
	text := MarshalNote(note)

	if text != "10123 [x] 09 Dec 20 21:05 +0000 test" {
		test.Fail()
	}
}

func TestMarshalNotes_empty(test *testing.T) {
	text := MarshalNotes([]todo.Note{})

	if text != "" {
		test.Fail()
	}
}

func TestMarshalNotes_nonEmpty(test *testing.T) {
	createdAt := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	notes := []todo.Note{
		todo.Note{ID: 100, CreatedAt: createdAt, IsDone: false, Text: "one"},
		todo.Note{
			ID:        101,
			CreatedAt: createdAt.Add(time.Hour),
			IsDone:    true,
			Text:      "two",
		},
		todo.Note{
			ID:        102,
			CreatedAt: createdAt.Add(2 * time.Hour),
			IsDone:    false,
			Text:      "three",
		},
		todo.Note{
			ID:        103,
			CreatedAt: createdAt.Add(3 * time.Hour),
			IsDone:    true,
			Text:      "four",
		},
		todo.Note{
			ID:        104,
			CreatedAt: createdAt.Add(4 * time.Hour),
			IsDone:    false,
			Text:      "five",
		},
	}
	text := MarshalNotes(notes)

	wantedText := "10100 [_] 09 Dec 20 21:05 +0000 one\n" +
		"10101 [x] 09 Dec 20 22:05 +0000 two\n" +
		"10102 [_] 09 Dec 20 23:05 +0000 three\n" +
		"10103 [x] 10 Dec 20 00:05 +0000 four\n" +
		"10104 [_] 10 Dec 20 01:05 +0000 five\n"
	if text != wantedText {
		test.Fail()
	}
}
