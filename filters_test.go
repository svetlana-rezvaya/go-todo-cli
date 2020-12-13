package main

import (
	"reflect"
	"testing"
	"time"
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

func Test_filterByTime_empty(test *testing.T) {
	minimum := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	maximum := minimum.Add(567 * time.Hour)
	notes := filterByTime([]note{}, minimum, maximum)

	wantedNotes := []note{}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func Test_filterByTime_nonEmpty(test *testing.T) {
	createdAt := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	minimum := createdAt.Add(2 * time.Hour)
	maximum := createdAt.Add(3 * time.Hour)
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
	notes = filterByTime(notes, minimum, maximum)

	wantedNotes := []note{
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
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}
