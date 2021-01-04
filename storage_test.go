package main

import (
	"io/ioutil"
	"testing"
	"time"
)

func Test_saveNotes_empty(test *testing.T) {
	file, err := ioutil.TempFile("", "test*.data")
	if err != nil {
		test.FailNow()
	}
	file.Close()

	err = saveNotes(file.Name(), []note{})
	if err != nil {
		test.FailNow()
	}

	textBytes, err := ioutil.ReadFile(file.Name())
	if err != nil {
		test.FailNow()
	}

	if len(textBytes) != 0 {
		test.Fail()
	}
}

func Test_saveNotes_nonEmpty(test *testing.T) {
	file, err := ioutil.TempFile("", "test*.data")
	if err != nil {
		test.FailNow()
	}
	file.Close()

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
	err = saveNotes(file.Name(), notes)
	if err != nil {
		test.FailNow()
	}

	textBytes, err := ioutil.ReadFile(file.Name())
	if err != nil {
		test.FailNow()
	}

	wantedText := "10100 [_] 09 Dec 20 21:05 +0000 one\n" +
		"10101 [x] 09 Dec 20 22:05 +0000 two\n" +
		"10102 [_] 09 Dec 20 23:05 +0000 three\n" +
		"10103 [x] 10 Dec 20 00:05 +0000 four\n" +
		"10104 [_] 10 Dec 20 01:05 +0000 five\n"
	if string(textBytes) != wantedText {
		test.Fail()
	}
}
