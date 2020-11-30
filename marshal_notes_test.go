package main

import "testing"

func Test_marshalNote_isNotDone(test *testing.T) {
	note := note{ID: 123, IsDone: false, Text: "test"}
	text := marshalNote(note)
	if text != "10123 [ ] test" {
		test.Fail()
	}
}

func Test_marshalNote_isDone(test *testing.T) {
	note := note{ID: 123, IsDone: true, Text: "test"}
	text := marshalNote(note)
	if text != "10123 [x] test" {
		test.Fail()
	}
}
