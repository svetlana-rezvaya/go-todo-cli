package cli

import (
	"reflect"
	"testing"
	"time"

	todo "github.com/svetlana-rezvaya/go-todo-cli"
)

func TestFilterByCommand_list(test *testing.T) {
	notes := []todo.Note{
		todo.Note{ID: 100, IsDone: false, Text: "one"},
		todo.Note{ID: 101, IsDone: true, Text: "two"},
		todo.Note{ID: 102, IsDone: false, Text: "three"},
		todo.Note{ID: 103, IsDone: true, Text: "four"},
		todo.Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := FilterByCommand(notes, "list")

	wantedNotes := []todo.Note{
		todo.Note{ID: 100, IsDone: false, Text: "one"},
		todo.Note{ID: 101, IsDone: true, Text: "two"},
		todo.Note{ID: 102, IsDone: false, Text: "three"},
		todo.Note{ID: 103, IsDone: true, Text: "four"},
		todo.Note{ID: 104, IsDone: false, Text: "five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestFilterByCommand_listDone(test *testing.T) {
	notes := []todo.Note{
		todo.Note{ID: 100, IsDone: false, Text: "one"},
		todo.Note{ID: 101, IsDone: true, Text: "two"},
		todo.Note{ID: 102, IsDone: false, Text: "three"},
		todo.Note{ID: 103, IsDone: true, Text: "four"},
		todo.Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := FilterByCommand(notes, "list done")

	wantedNotes := []todo.Note{
		todo.Note{ID: 101, IsDone: true, Text: "two"},
		todo.Note{ID: 103, IsDone: true, Text: "four"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestFilterByCommand_listToDo(test *testing.T) {
	notes := []todo.Note{
		todo.Note{ID: 100, IsDone: false, Text: "one"},
		todo.Note{ID: 101, IsDone: true, Text: "two"},
		todo.Note{ID: 102, IsDone: false, Text: "three"},
		todo.Note{ID: 103, IsDone: true, Text: "four"},
		todo.Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := FilterByCommand(notes, "list to do")

	wantedNotes := []todo.Note{
		todo.Note{ID: 100, IsDone: false, Text: "one"},
		todo.Note{ID: 102, IsDone: false, Text: "three"},
		todo.Note{ID: 104, IsDone: false, Text: "five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestFilterByCommand_listUnknown(test *testing.T) {
	notes := []todo.Note{
		todo.Note{ID: 100, IsDone: false, Text: "one"},
		todo.Note{ID: 101, IsDone: true, Text: "two"},
		todo.Note{ID: 102, IsDone: false, Text: "three"},
		todo.Note{ID: 103, IsDone: true, Text: "four"},
		todo.Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := FilterByCommand(notes, "list unknown")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "unknown parameter for 'list' command: unknown"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func TestFilterByCommand_find(test *testing.T) {
	notes := []todo.Note{
		todo.Note{ID: 100, IsDone: false, Text: "task No.1, one"},
		todo.Note{ID: 101, IsDone: true, Text: "task No.1, two"},
		todo.Note{ID: 102, IsDone: false, Text: "task No.2, three"},
		todo.Note{ID: 103, IsDone: true, Text: "task No.2, four"},
		todo.Note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	notes, err := FilterByCommand(notes, "find No.2")

	wantedNotes := []todo.Note{
		todo.Note{ID: 102, IsDone: false, Text: "task No.2, three"},
		todo.Note{ID: 103, IsDone: true, Text: "task No.2, four"},
		todo.Note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestFilterByCommand_findWithoutParameter(test *testing.T) {
	notes := []todo.Note{
		todo.Note{ID: 100, IsDone: false, Text: "task No.1, one"},
		todo.Note{ID: 101, IsDone: true, Text: "task No.1, two"},
		todo.Note{ID: 102, IsDone: false, Text: "task No.2, three"},
		todo.Note{ID: 103, IsDone: true, Text: "task No.2, four"},
		todo.Note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	notes, err := FilterByCommand(notes, "find")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "query missing in 'find' command"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func TestFilterByCommand_date(test *testing.T) {
	createdAt := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	notes := []todo.Note{
		todo.Note{ID: 100, CreatedAt: createdAt, IsDone: false, Text: "one"},
		todo.Note{
			ID:        101,
			CreatedAt: createdAt.Add(time.Hour),
			IsDone:    true,
			Text:      "two",
		},
		todo.Note{
			ID:        102,
			CreatedAt: createdAt.Add(2 * time.Hour),
			IsDone:    false,
			Text:      "three",
		},
		todo.Note{
			ID:        103,
			CreatedAt: createdAt.Add(3 * time.Hour),
			IsDone:    true,
			Text:      "four",
		},
		todo.Note{
			ID:        104,
			CreatedAt: createdAt.Add(4 * time.Hour),
			IsDone:    false,
			Text:      "five",
		},
	}
	notes, err := FilterByCommand(notes, "date 09 Dec 20")

	wantedNotes := []todo.Note{
		todo.Note{ID: 100, CreatedAt: createdAt, IsDone: false, Text: "one"},
		todo.Note{
			ID:        101,
			CreatedAt: createdAt.Add(time.Hour),
			IsDone:    true,
			Text:      "two",
		},
		todo.Note{
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

func TestFilterByCommand_dateWithoutParameter(test *testing.T) {
	createdAt := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	notes := []todo.Note{
		todo.Note{ID: 100, CreatedAt: createdAt, IsDone: false, Text: "one"},
		todo.Note{
			ID:        101,
			CreatedAt: createdAt.Add(time.Hour),
			IsDone:    true,
			Text:      "two",
		},
		todo.Note{
			ID:        102,
			CreatedAt: createdAt.Add(2 * time.Hour),
			IsDone:    false,
			Text:      "three",
		},
		todo.Note{
			ID:        103,
			CreatedAt: createdAt.Add(3 * time.Hour),
			IsDone:    true,
			Text:      "four",
		},
		todo.Note{
			ID:        104,
			CreatedAt: createdAt.Add(4 * time.Hour),
			IsDone:    false,
			Text:      "five",
		},
	}
	notes, err := FilterByCommand(notes, "date")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "parameter missing in 'date' command"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func TestFilterByCommand_dateWithIncorrectParameter(test *testing.T) {
	createdAt := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	notes := []todo.Note{
		todo.Note{ID: 100, CreatedAt: createdAt, IsDone: false, Text: "one"},
		todo.Note{
			ID:        101,
			CreatedAt: createdAt.Add(time.Hour),
			IsDone:    true,
			Text:      "two",
		},
		todo.Note{
			ID:        102,
			CreatedAt: createdAt.Add(2 * time.Hour),
			IsDone:    false,
			Text:      "three",
		},
		todo.Note{
			ID:        103,
			CreatedAt: createdAt.Add(3 * time.Hour),
			IsDone:    true,
			Text:      "four",
		},
		todo.Note{
			ID:        104,
			CreatedAt: createdAt.Add(4 * time.Hour),
			IsDone:    false,
			Text:      "five",
		},
	}
	notes, err := FilterByCommand(notes, "date 09 12 20")

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

func TestFilterByCommand_unknown(test *testing.T) {
	notes := []todo.Note{
		todo.Note{ID: 100, IsDone: false, Text: "one"},
		todo.Note{ID: 101, IsDone: true, Text: "two"},
		todo.Note{ID: 102, IsDone: false, Text: "three"},
		todo.Note{ID: 103, IsDone: true, Text: "four"},
		todo.Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := FilterByCommand(notes, "unknown")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "unknown command: unknown"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func TestFilterByMultiCommand_withOneCommand(test *testing.T) {
	notes := []todo.Note{
		todo.Note{ID: 100, IsDone: false, Text: "task No.1, one"},
		todo.Note{ID: 101, IsDone: true, Text: "task No.1, two"},
		todo.Note{ID: 102, IsDone: false, Text: "task No.2, three"},
		todo.Note{ID: 103, IsDone: true, Text: "task No.2, four"},
		todo.Note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	notes, err := FilterByMultiCommand(notes, "find No.2")

	wantedNotes := []todo.Note{
		todo.Note{ID: 102, IsDone: false, Text: "task No.2, three"},
		todo.Note{ID: 103, IsDone: true, Text: "task No.2, four"},
		todo.Note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestFilterByMultiCommand_withFewCommands(test *testing.T) {
	notes := []todo.Note{
		todo.Note{ID: 100, IsDone: false, Text: "task No.1, one"},
		todo.Note{ID: 101, IsDone: true, Text: "task No.1, two"},
		todo.Note{ID: 102, IsDone: false, Text: "task No.2, three"},
		todo.Note{ID: 103, IsDone: true, Text: "task No.2, four"},
		todo.Note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	notes, err := FilterByMultiCommand(notes, "find No.2 | list to do")

	wantedNotes := []todo.Note{
		todo.Note{ID: 102, IsDone: false, Text: "task No.2, three"},
		todo.Note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestFilterByMultiCommand_withEmptyCommand(test *testing.T) {
	notes := []todo.Note{
		todo.Note{ID: 100, IsDone: false, Text: "task No.1, one"},
		todo.Note{ID: 101, IsDone: true, Text: "task No.1, two"},
		todo.Note{ID: 102, IsDone: false, Text: "task No.2, three"},
		todo.Note{ID: 103, IsDone: true, Text: "task No.2, four"},
		todo.Note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	notes, err := FilterByMultiCommand(notes, "find No.2 | | list to do")

	wantedNotes := []todo.Note{
		todo.Note{ID: 102, IsDone: false, Text: "task No.2, three"},
		todo.Note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestFilterByMultiCommand_error(test *testing.T) {
	notes := []todo.Note{
		todo.Note{ID: 100, IsDone: false, Text: "task No.1, one"},
		todo.Note{ID: 101, IsDone: true, Text: "task No.1, two"},
		todo.Note{ID: 102, IsDone: false, Text: "task No.2, three"},
		todo.Note{ID: 103, IsDone: true, Text: "task No.2, four"},
		todo.Note{ID: 104, IsDone: false, Text: "task No.2, five"},
	}
	notes, err := FilterByMultiCommand(notes, "find No.2 | list unknown")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "unable to filter by command #2: " +
		"unknown parameter for 'list' command: unknown"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}
