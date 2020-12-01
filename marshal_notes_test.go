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

func Test_marshalNotes_empty(test *testing.T) {
	text := marshalNotes([]note{})

	if text != "" {
		test.Fail()
	}
}

func Test_marshalNotes_nonEmpty(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	text := marshalNotes(notes)

	wantedText := "10100 [ ] one\n" +
		"10101 [x] two\n" +
		"10102 [ ] three\n" +
		"10103 [x] four\n" +
		"10104 [ ] five\n"
	if text != wantedText {
		test.Fail()
	}
}
