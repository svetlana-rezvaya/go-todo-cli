package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_unmarshalNote_successAndIsNotDone(test *testing.T) {
	noteObject, err :=
		unmarshalNote("10123 [_] 09 Dec 20 21:05 +0000 one two three")
	noteObject.CreatedAt = noteObject.CreatedAt.In(time.UTC)

	createdAt := time.Date(2020, time.December, 9, 21, 5, 0, 0, time.UTC)
	wantedNoteObject := note{
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

func Test_unmarshalNote_successAndIsDone(test *testing.T) {
	noteObject, err :=
		unmarshalNote("10123 [x] 09 Dec 20 21:05 +0000 one two three")
	noteObject.CreatedAt = noteObject.CreatedAt.In(time.UTC)

	createdAt := time.Date(2020, time.December, 9, 21, 5, 0, 0, time.UTC)
	wantedNoteObject := note{
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

func Test_unmarshalNote_errorWithPartCount(test *testing.T) {
	noteObject, err := unmarshalNote("10123 [_] 09 Dec 20")

	wantedNoteObject := note{}
	if !reflect.DeepEqual(noteObject, wantedNoteObject) {
		test.Fail()
	}

	if err == nil || err.Error() != "not enough parts in the line" {
		test.Fail()
	}
}

func Test_unmarshalNote_errorWithIncorrectID(test *testing.T) {
	noteObject, err :=
		unmarshalNote("incorrect [_] 09 Dec 20 21:05 +0000 one two three")

	wantedNoteObject := note{}
	if !reflect.DeepEqual(noteObject, wantedNoteObject) {
		test.Fail()
	}

	wantedErrStr := "unable to parse the note ID: " +
		"strconv.Atoi: parsing \"incorrect\": invalid syntax"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func Test_unmarshalNote_errorWithIncorrectTimestamp(test *testing.T) {
	noteObject, err :=
		unmarshalNote("10123 [_] 09 12 20 21:05 +0000 one two three")

	wantedNoteObject := note{}
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

func Test_unmarshalNotes_empty(test *testing.T) {
	notes, err := unmarshalNotes("")

	wantedNotes := []note{}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_unmarshalNotes_success(test *testing.T) {
	text := "10100 [_] 09 Dec 20 21:05 +0000 one\n" +
		"10101 [x] 09 Dec 20 22:05 +0000 two\n" +
		"\n" +
		"10102 [_] 09 Dec 20 23:05 +0000 three\n" +
		"10103 [x] 10 Dec 20 00:05 +0000 four\n" +
		"\n" +
		"10104 [_] 10 Dec 20 01:05 +0000 five\n"
	notes, err := unmarshalNotes(text)
	for index := range notes {
		notes[index].CreatedAt = notes[index].CreatedAt.In(time.UTC)
	}

	createdAt := time.Date(2020, time.December, 9, 21, 5, 0, 0, time.UTC)
	wantedNotes := []note{
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
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}
