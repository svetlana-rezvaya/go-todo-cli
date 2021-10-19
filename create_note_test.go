package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCreateNote_empty(test *testing.T) {
	notes := CreateNote([]Note{}, "test")

	if len(notes) != 0 {
		if time.Since(notes[len(notes)-1].CreatedAt) > time.Minute {
			test.Fail()
		}

		notes[len(notes)-1].CreatedAt = time.Time{}
	}

	wantedNotes := []Note{
		Note{ID: 1, IsDone: false, Text: "test"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func TestCreateNote_nonEmpty(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes = CreateNote(notes, "test")

	if len(notes) != 0 {
		if time.Since(notes[len(notes)-1].CreatedAt) > time.Minute {
			test.Fail()
		}

		notes[len(notes)-1].CreatedAt = time.Time{}
	}

	wantedNotes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
		Note{ID: 105, IsDone: false, Text: "test"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func Test_getMaximalID_empty(test *testing.T) {
	maximalID := getMaximalID([]Note{})

	if maximalID != 0 {
		test.Fail()
	}
}

func Test_getMaximalID_nonEmpty(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	maximalID := getMaximalID(notes)

	if maximalID != 104 {
		test.Fail()
	}
}
