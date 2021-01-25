package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_getMaximalID_empty(test *testing.T) {
	maximalID := getMaximalID([]note{})

	if maximalID != 0 {
		test.Fail()
	}
}

func Test_getMaximalID_nonEmpty(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	maximalID := getMaximalID(notes)

	if maximalID != 104 {
		test.Fail()
	}
}

func Test_createNote_empty(test *testing.T) {
	notes := createNote([]note{}, "test")

	if len(notes) != 0 {
		if time.Since(notes[len(notes)-1].CreatedAt) > time.Minute {
			test.Fail()
		}

		notes[len(notes)-1].CreatedAt = time.Time{}
	}

	wantedNotes := []note{
		note{ID: 1, IsDone: false, Text: "test"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func Test_createNote_nonEmpty(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes = createNote(notes, "test")

	if len(notes) != 0 {
		if time.Since(notes[len(notes)-1].CreatedAt) > time.Minute {
			test.Fail()
		}

		notes[len(notes)-1].CreatedAt = time.Time{}
	}

	wantedNotes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
		note{ID: 105, IsDone: false, Text: "test"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}
