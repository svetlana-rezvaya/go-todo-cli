package main

import (
	"reflect"
	"testing"
)

func Test_deleteNote_empty(test *testing.T) {
	notes := deleteNote([]note{}, 23)

	wantedNotes := []note{}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func Test_deleteNote_nonEmptyAndNonExistingID(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes = deleteNote(notes, 23)

	wantedNotes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func Test_deleteNote_nonEmptyAndExistingID(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes = deleteNote(notes, 102)

	wantedNotes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}
