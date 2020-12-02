package main

import "testing"

func Test_getMaximalID_empty(test *testing.T) {
	maximalID := getMaximalID([]note{})

	if maximalID != 0 {
		test.Fail()
	}
}

func Test_getMaximalID_nonEmpty(test *testing.T) {
	notes := []note{
		note{ID: 100, IsDone: false, Text: "one"},
		note{ID: 101, IsDone: true, Text: "two"},
		note{ID: 102, IsDone: false, Text: "three"},
		note{ID: 103, IsDone: true, Text: "four"},
		note{ID: 104, IsDone: false, Text: "five"},
	}
	maximalID := getMaximalID(notes)

	if maximalID != 104 {
		test.Fail()
	}
}
