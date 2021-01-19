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

func Test_filterByDate_empty(test *testing.T) {
	date := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	notes := filterByDate([]note{}, date)

	wantedNotes := []note{}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func Test_filterByDate_nonEmpty(test *testing.T) {
	createdAt := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	date := createdAt.Add(2 * time.Hour)
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
	notes = filterByDate(notes, date)

	wantedNotes := []note{
		note{ID: 100, CreatedAt: createdAt, IsDone: false, Text: "one"},
		note{ID: 101, CreatedAt: createdAt.Add(time.Hour), IsDone: true, Text: "two"},
		note{
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

func Test_filterByCommand_list(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := filterByCommand(notes, "list")

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

	if err != nil {
		test.Fail()
	}
}

func Test_filterByCommand_listDone(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := filterByCommand(notes, "list done")

	wantedNotes := []note{
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 103, IsDone: true, Text: "four"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_filterByCommand_listToDo(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := filterByCommand(notes, "list to do")

	wantedNotes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_filterByCommand_listUnknown(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := filterByCommand(notes, "list unknown")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "unknown parameter for 'list' command: unknown"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func Test_filterByCommand_find(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "task No.1, one"},
		note{ID: 101, IsDone: true, Text: "task No.1, two"},
		note{ID: 102, IsDone: false, Text: "task No.2, three"},
		note{ID: 103, IsDone: true, Text: "task No.2, four"},
		note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	notes, err := filterByCommand(notes, "find No.2")

	wantedNotes := []note{
		note{ID: 102, IsDone: false, Text: "task No.2, three"},
		note{ID: 103, IsDone: true, Text: "task No.2, four"},
		note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_filterByCommand_findWithoutParameter(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "task No.1, one"},
		note{ID: 101, IsDone: true, Text: "task No.1, two"},
		note{ID: 102, IsDone: false, Text: "task No.2, three"},
		note{ID: 103, IsDone: true, Text: "task No.2, four"},
		note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	notes, err := filterByCommand(notes, "find")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "query missing in 'find' command"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func Test_filterByCommand_date(test *testing.T) {
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
	notes, err := filterByCommand(notes, "date 09 Dec 20")

	wantedNotes := []note{
		note{ID: 100, CreatedAt: createdAt, IsDone: false, Text: "one"},
		note{ID: 101, CreatedAt: createdAt.Add(time.Hour), IsDone: true, Text: "two"},
		note{
			ID:        102,
			CreatedAt: createdAt.Add(2 * time.Hour),
			IsDone:    false,
			Text:      "three",
		},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_filterByCommand_dateWithoutParameter(test *testing.T) {
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
	notes, err := filterByCommand(notes, "date")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "parameter missing in 'date' command"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func Test_filterByCommand_dateWithIncorrectParameter(test *testing.T) {
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
	notes, err := filterByCommand(notes, "date 09 12 20")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "unable to parse the 'date' command parameter: " +
		"parsing time \"09 12 20\" as \"02 Jan 06\": " +
		"cannot parse \"12 20\" as \"Jan\""
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func Test_filterByCommand_unknown(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := filterByCommand(notes, "unknown")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "unknown command: unknown"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func Test_filterByMultiCommand_withOneCommand(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "task No.1, one"},
		note{ID: 101, IsDone: true, Text: "task No.1, two"},
		note{ID: 102, IsDone: false, Text: "task No.2, three"},
		note{ID: 103, IsDone: true, Text: "task No.2, four"},
		note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	notes, err := filterByMultiCommand(notes, "find No.2")

	wantedNotes := []note{
		note{ID: 102, IsDone: false, Text: "task No.2, three"},
		note{ID: 103, IsDone: true, Text: "task No.2, four"},
		note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_filterByMultiCommand_withFewCommands(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "task No.1, one"},
		note{ID: 101, IsDone: true, Text: "task No.1, two"},
		note{ID: 102, IsDone: false, Text: "task No.2, three"},
		note{ID: 103, IsDone: true, Text: "task No.2, four"},
		note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	notes, err := filterByMultiCommand(notes, "find No.2 | list to do")

	wantedNotes := []note{
		note{ID: 102, IsDone: false, Text: "task No.2, three"},
		note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_filterByMultiCommand_withEmptyCommand(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "task No.1, one"},
		note{ID: 101, IsDone: true, Text: "task No.1, two"},
		note{ID: 102, IsDone: false, Text: "task No.2, three"},
		note{ID: 103, IsDone: true, Text: "task No.2, four"},
		note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	notes, err := filterByMultiCommand(notes, "find No.2 | | list to do")

	wantedNotes := []note{
		note{ID: 102, IsDone: false, Text: "task No.2, three"},
		note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_filterByMultiCommand_error(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "task No.1, one"},
		note{ID: 101, IsDone: true, Text: "task No.1, two"},
		note{ID: 102, IsDone: false, Text: "task No.2, three"},
		note{ID: 103, IsDone: true, Text: "task No.2, four"},
		note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	notes, err := filterByMultiCommand(notes, "find No.2 | list unknown")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "unable to filter by command #2: " +
		"unknown parameter for 'list' command: unknown"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}
