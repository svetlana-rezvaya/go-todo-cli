package encoding

import (
	"reflect"
	"testing"
	"time"

	todo "github.com/svetlana-rezvaya/go-todo-cli"
)

func TestUnmarshalNote_successAndIsNotDone(test *testing.T) {
	noteObject, err :=
		UnmarshalNote("10123 [_] 09 Dec 20 21:05 +0000 one two three")
	noteObject.CreatedAt = noteObject.CreatedAt.In(time.UTC)

	createdAt := time.Date(2020, time.December, 9, 21, 5, 0, 0, time.UTC)
	wantedNoteObject := todo.Note{
		ID:        123,
		CreatedAt: createdAt,
		IsDone:    false,
		Text:      "one two three",
	}
	if !reflect.DeepEqual(noteObject, wantedNoteObject) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestUnmarshalNote_successAndIsDone(test *testing.T) {
	noteObject, err :=
		UnmarshalNote("10123 [x] 09 Dec 20 21:05 +0000 one two three")
	noteObject.CreatedAt = noteObject.CreatedAt.In(time.UTC)

	createdAt := time.Date(2020, time.December, 9, 21, 5, 0, 0, time.UTC)
	wantedNoteObject := todo.Note{
		ID:        123,
		CreatedAt: createdAt,
		IsDone:    true,
		Text:      "one two three",
	}
	if !reflect.DeepEqual(noteObject, wantedNoteObject) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestUnmarshalNote_errorWithPartCount(test *testing.T) {
	noteObject, err := UnmarshalNote("10123 [_] 09 Dec 20")

	wantedNoteObject := todo.Note{}
	if !reflect.DeepEqual(noteObject, wantedNoteObject) {
		test.Fail()
	}

	if err == nil || err.Error() != "not enough parts in the line" {
		test.Fail()
	}
}

func TestUnmarshalNote_errorWithIncorrectID(test *testing.T) {
	noteObject, err :=
		UnmarshalNote("incorrect [_] 09 Dec 20 21:05 +0000 one two three")

	wantedNoteObject := todo.Note{}
	if !reflect.DeepEqual(noteObject, wantedNoteObject) {
		test.Fail()
	}

	wantedErrStr := "unable to parse the note ID: " +
		"strconv.Atoi: parsing \"incorrect\": invalid syntax"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func TestUnmarshalNote_errorWithIncorrectTimestamp(test *testing.T) {
	noteObject, err :=
		UnmarshalNote("10123 [_] 09 12 20 21:05 +0000 one two three")

	wantedNoteObject := todo.Note{}
	if !reflect.DeepEqual(noteObject, wantedNoteObject) {
		test.Fail()
	}

	wantedErrStr := "unable to parse the note creation timestamp: " +
		"parsing time \"09 12 20 21:05 +0000\" as \"02 Jan 06 15:04 -0700\": " +
		"cannot parse \"12 20 21:05 +0000\" as \"Jan\""
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func TestUnmarshalNotes_empty(test *testing.T) {
	notes, err := UnmarshalNotes("")

	wantedNotes := []todo.Note{}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestUnmarshalNotes_success(test *testing.T) {
	text := "10100 [_] 09 Dec 20 21:05 +0000 one\n" +
		"10101 [x] 09 Dec 20 22:05 +0000 two\n" +
		"\n" +
		"10102 [_] 09 Dec 20 23:05 +0000 three\n" +
		"10103 [x] 10 Dec 20 00:05 +0000 four\n" +
		"\n" +
		"10104 [_] 10 Dec 20 01:05 +0000 five\n"
	notes, err := UnmarshalNotes(text)
	for index := range notes {
		notes[index].CreatedAt = notes[index].CreatedAt.In(time.UTC)
	}

	createdAt := time.Date(2020, time.December, 9, 21, 5, 0, 0, time.UTC)
	wantedNotes := []todo.Note{
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
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestUnmarshalNotes_error(test *testing.T) {
	text := "10100 [_] 09 Dec 20 21:05 +0000 one\n" +
		"10101 [x] 09 Dec 20 22:05 +0000 two\n" +
		"\n" +
		"10102 [_] 09 Dec 20 23:05 +0000 three\n" +
		"incorrect [x] 10 Dec 20 00:05 +0000 four\n" +
		"\n" +
		"10104 [_] 10 Dec 20 01:05 +0000 five\n"
	notes, err := UnmarshalNotes(text)

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "unable to unmarshal the note in line #5: " +
		"unable to parse the note ID: " +
		"strconv.Atoi: parsing \"incorrect\": invalid syntax"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}
