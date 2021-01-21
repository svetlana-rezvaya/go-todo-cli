package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_updateStatus_empty(test *testing.T) {
	notes := []note{}
	updateStatus(notes, 23, true)

	wantedNotes := []note{}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func Test_updateStatus_nonEmptyAndNonExistingID(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	updateStatus(notes, 23, true)

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

func Test_updateStatus_nonEmptyAndExistingID(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	updateStatus(notes, 102, true)

	if len(notes) >= 3 {
		if time.Since(notes[2].UpdatedAt) > time.Minute {
			test.Fail()
		}

		notes[2].UpdatedAt = time.Time{}
	}

	wantedNotes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: true, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func Test_updateUsingCommand_add(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := updateUsingCommand(notes, "add test")

	if len(notes) != 0 {
		if time.Since(notes[len(notes)-1].CreatedAt) > time.Minute {
			test.Fail()
		}
		if time.Since(notes[len(notes)-1].UpdatedAt) > time.Minute {
			test.Fail()
		}

		notes[len(notes)-1].CreatedAt = time.Time{}
		notes[len(notes)-1].UpdatedAt = time.Time{}
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

	if err != nil {
		test.Fail()
	}
}

func Test_updateUsingCommand_addWithoutText(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := updateUsingCommand(notes, "add")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "text missing in 'add' command"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func Test_updateUsingCommand_check(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := updateUsingCommand(notes, "check 10102")

	if len(notes) >= 3 {
		if time.Since(notes[2].UpdatedAt) > time.Minute {
			test.Fail()
		}

		notes[2].UpdatedAt = time.Time{}
	}

	wantedNotes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: true, Text: "three"},
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

func Test_updateUsingCommand_checkWithoutID(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := updateUsingCommand(notes, "check")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "unable to get the note ID: note ID missing"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func Test_updateUsingCommand_uncheck(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := updateUsingCommand(notes, "uncheck 10101")

	if len(notes) >= 2 {
		if time.Since(notes[1].UpdatedAt) > time.Minute {
			test.Fail()
		}

		notes[1].UpdatedAt = time.Time{}
	}

	wantedNotes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: false, Text: "two"},
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

func Test_updateUsingCommand_uncheckWithoutID(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := updateUsingCommand(notes, "uncheck")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "unable to get the note ID: note ID missing"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func Test_updateUsingCommand_delete(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := updateUsingCommand(notes, "delete 10102")

	wantedNotes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
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

func Test_updateUsingCommand_deleteWithoutID(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := updateUsingCommand(notes, "delete")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "unable to get the note ID: note ID missing"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}
