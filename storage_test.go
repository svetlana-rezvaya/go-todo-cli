package todo

import (
	"io/ioutil"
	"reflect"
	"testing"
	"time"
)

func TestLoadNotes_empty(test *testing.T) {
	file, err := ioutil.TempFile("", "test*.data")
	if err != nil {
		test.FailNow()
	}
	file.Close()

	notes, err := LoadNotes(file.Name())

	wantedNotes := []Note{}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestLoadNotes_nonEmpty(test *testing.T) {
	file, err := ioutil.TempFile("", "test*.data")
	if err != nil {
		test.FailNow()
	}
	file.Close()

	text := "10100 [_] 09 Dec 20 21:05 +0000 one\n" +
		"10101 [x] 09 Dec 20 22:05 +0000 two\n" +
		"10102 [_] 09 Dec 20 23:05 +0000 three\n" +
		"10103 [x] 10 Dec 20 00:05 +0000 four\n" +
		"10104 [_] 10 Dec 20 01:05 +0000 five\n"
	err = ioutil.WriteFile(file.Name(), []byte(text), 0)
	if err != nil {
		test.FailNow()
	}

	notes, err := LoadNotes(file.Name())
	for index := range notes {
		notes[index].CreatedAt = notes[index].CreatedAt.In(time.UTC)
	}

	createdAt := time.Date(2020, time.December, 9, 21, 5, 0, 0, time.UTC)
	wantedNotes := []Note{
		Note{ID: 100, CreatedAt: createdAt, IsDone: false, Text: "one"},
		Note{ID: 101, CreatedAt: createdAt.Add(time.Hour), IsDone: true, Text: "two"},
		Note{
			ID:        102,
			CreatedAt: createdAt.Add(2 * time.Hour),
			IsDone:    false,
			Text:      "three",
		},
		Note{
			ID:        103,
			CreatedAt: createdAt.Add(3 * time.Hour),
			IsDone:    true,
			Text:      "four",
		},
		Note{
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

func TestLoadNotes_error(test *testing.T) {
	file, err := ioutil.TempFile("", "test*.data")
	if err != nil {
		test.FailNow()
	}
	file.Close()

	text := "10100 [_] 09 Dec 20 21:05 +0000 one\n" +
		"10101 [x] 09 Dec 20 22:05 +0000 two\n" +
		"10102 [_] 09 Dec 20 23:05 +0000 three\n" +
		"incorrect [x] 10 Dec 20 00:05 +0000 four\n" +
		"10104 [_] 10 Dec 20 01:05 +0000 five\n"
	err = ioutil.WriteFile(file.Name(), []byte(text), 0)
	if err != nil {
		test.FailNow()
	}

	notes, err := LoadNotes(file.Name())

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "unable to unmarshal a storage file: " +
		"unable to unmarshal the note in line #4: " +
		"unable to parse the note ID: " +
		"strconv.Atoi: parsing \"incorrect\": invalid syntax"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func TestSaveNotes_empty(test *testing.T) {
	file, err := ioutil.TempFile("", "test*.data")
	if err != nil {
		test.FailNow()
	}
	file.Close()

	err = SaveNotes(file.Name(), []Note{})
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

func TestSaveNotes_nonEmpty(test *testing.T) {
	file, err := ioutil.TempFile("", "test*.data")
	if err != nil {
		test.FailNow()
	}
	file.Close()

	createdAt := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	notes := []Note{
		Note{ID: 100, CreatedAt: createdAt, IsDone: false, Text: "one"},
		Note{ID: 101, CreatedAt: createdAt.Add(time.Hour), IsDone: true, Text: "two"},
		Note{
			ID:        102,
			CreatedAt: createdAt.Add(2 * time.Hour),
			IsDone:    false,
			Text:      "three",
		},
		Note{
			ID:        103,
			CreatedAt: createdAt.Add(3 * time.Hour),
			IsDone:    true,
			Text:      "four",
		},
		Note{
			ID:        104,
			CreatedAt: createdAt.Add(4 * time.Hour),
			IsDone:    false,
			Text:      "five",
		},
	}
	err = SaveNotes(file.Name(), notes)
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
