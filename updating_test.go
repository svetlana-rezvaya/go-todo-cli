package main

import (
	"reflect"
	"testing"
	"time"
)

func TestUpdateStatus_empty(test *testing.T) {
	notes := []Note{}
	UpdateStatus(notes, 23, true)

	wantedNotes := []Note{}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func TestUpdateStatus_nonEmptyAndNonExistingID(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	UpdateStatus(notes, 23, true)

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

func TestUpdateStatus_nonEmptyAndExistingID(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	UpdateStatus(notes, 102, true)

	wantedNotes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: true, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}
}

func TestUpdateUsingCommand_add(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := UpdateUsingCommand(notes, "add test")

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

	if err != nil {
		test.Fail()
	}
}

func TestUpdateUsingCommand_addWithoutText(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := UpdateUsingCommand(notes, "add")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "text missing in 'add' command"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func TestUpdateUsingCommand_check(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := UpdateUsingCommand(notes, "check 10102")

	wantedNotes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: true, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestUpdateUsingCommand_checkWithoutID(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := UpdateUsingCommand(notes, "check")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "unable to get the note ID: note ID missing"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func TestUpdateUsingCommand_uncheck(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := UpdateUsingCommand(notes, "uncheck 10101")

	wantedNotes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: false, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestUpdateUsingCommand_uncheckWithoutID(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := UpdateUsingCommand(notes, "uncheck")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "unable to get the note ID: note ID missing"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func TestUpdateUsingCommand_delete(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := UpdateUsingCommand(notes, "delete 10102")

	wantedNotes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	if !reflect.DeepEqual(notes, wantedNotes) {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestUpdateUsingCommand_deleteWithoutID(test *testing.T) {
	notes := []Note{
		Note{ID: 100, IsDone: false, Text: "one"},
		Note{ID: 101, IsDone: true, Text: "two"},
		Note{ID: 102, IsDone: false, Text: "three"},
		Note{ID: 103, IsDone: true, Text: "four"},
		Note{ID: 104, IsDone: false, Text: "five"},
	}
	notes, err := UpdateUsingCommand(notes, "delete")

	if notes != nil {
		test.Fail()
	}

	wantedErrStr := "unable to get the note ID: note ID missing"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}
