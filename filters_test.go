package todo

import (
	"reflect"
	"testing"
	"time"
)

func TestFilterByStatus_empty(test *testing.T) {
	notes := FilterByStatus([]Note{}, true)

	wantedNotes := []Note{}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func TestFilterByStatus_isDone(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes = FilterByStatus(notes, true)

	wantedNotes := []Note{
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 103, IsDone: true, Text: "four"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func TestFilterByStatus_isNotDone(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes = FilterByStatus(notes, false)

	wantedNotes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func TestFilterByText_empty(test *testing.T) {
	notes := FilterByText([]Note{}, "No.2")

	wantedNotes := []Note{}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func TestFilterByText_nonEmpty(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "task No.1, one"},
		Note{ID: 101, IsDone: true, Text: "task No.1, two"},
		Note{ID: 102, IsDone: false, Text: "task No.2, three"},
		Note{ID: 103, IsDone: true, Text: "task No.2, four"},
		Note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	notes = FilterByText(notes, "No.2")

	wantedNotes := []Note{
		Note{ID: 102, IsDone: false, Text: "task No.2, three"},
		Note{ID: 103, IsDone: true, Text: "task No.2, four"},
		Note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func TestFilterByTime_empty(test *testing.T) {
	minimum := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	maximum := minimum.Add(567 * time.Hour)
	notes := FilterByTime([]Note{}, minimum, maximum)

	wantedNotes := []Note{}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func TestFilterByTime_nonEmpty(test *testing.T) {
	createdAt := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	minimum := createdAt.Add(2 * time.Hour)
	maximum := createdAt.Add(3 * time.Hour)
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
	notes = FilterByTime(notes, minimum, maximum)

	wantedNotes := []Note{
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
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func TestFilterByDate_empty(test *testing.T) {
	date := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	notes := FilterByDate([]Note{}, date)

	wantedNotes := []Note{}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func TestFilterByDate_nonEmpty(test *testing.T) {
	createdAt := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	date := createdAt.Add(2 * time.Hour)
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
	notes = FilterByDate(notes, date)

	wantedNotes := []Note{
		Note{ID: 100, CreatedAt: createdAt, IsDone: false, Text: "one"},
		Note{ID: 101, CreatedAt: createdAt.Add(time.Hour), IsDone: true, Text: "two"},
		Note{
			ID:        102,
			CreatedAt: createdAt.Add(2 * time.Hour),
			IsDone:    false,
			Text:      "three",
		},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}
