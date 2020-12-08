package main

import (
	"reflect"
	"testing"
)

func Test_filterByStatus_empty(test *testing.T) {
	notes := filterByStatus([]note{}, true)

	wantedNotes := []note{}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func Test_filterByStatus_isDone(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes = filterByStatus(notes, true)

	wantedNotes := []note{
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 103, IsDone: true, Text: "four"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func Test_filterByStatus_isNotDone(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes = filterByStatus(notes, false)

	wantedNotes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func Test_filterByText_empty(test *testing.T) {
	notes := filterByText([]note{}, "No.2")

	wantedNotes := []note{}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func Test_filterByText_nonEmpty(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "task No.1, one"},
		note{ID: 101, IsDone: true, Text: "task No.1, two"},
		note{ID: 102, IsDone: false, Text: "task No.2, three"},
		note{ID: 103, IsDone: true, Text: "task No.2, four"},
		note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	notes = filterByText(notes, "No.2")

	wantedNotes := []note{
		note{ID: 102, IsDone: false, Text: "task No.2, three"},
		note{ID: 103, IsDone: true, Text: "task No.2, four"},
		note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}
