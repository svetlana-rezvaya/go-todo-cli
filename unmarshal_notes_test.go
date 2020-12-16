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
