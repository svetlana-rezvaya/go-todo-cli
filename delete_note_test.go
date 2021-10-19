package main

import (
	"reflect"
	"testing"
)

func TestDeleteNote_empty(test *testing.T) {
	notes := DeleteNote([]Note{}, 23)

	wantedNotes := []Note{}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func TestDeleteNote_nonEmptyAndNonExistingID(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes = DeleteNote(notes, 23)

	wantedNotes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func TestDeleteNote_nonEmptyAndExistingID(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes = DeleteNote(notes, 102)

	wantedNotes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}
