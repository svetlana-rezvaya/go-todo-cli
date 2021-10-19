package todo

import (
	"testing"
)

func Test_getParameter(test *testing.T) {
	parameter := getParameter("add test", "add")

	if parameter != "test" {
		test.Fail()
	}
}

func Test_getIDParameter_success(test *testing.T) {
	id, err := getIDParameter("delete 10101", "delete")

	if id != 101 {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_getIDParameter_errorWithoutID(test *testing.T) {
	id, err := getIDParameter("delete", "delete")

	if id != 0 {
		test.Fail()
	}

	wantedErrStr := "note ID missing"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func Test_getIDParameter_errorWithIncorrectID(test *testing.T) {
	id, err := getIDParameter("delete incorrect", "delete")

	if id != 0 {
		test.Fail()
	}

	wantedErrStr := "unable to parse the note ID: " +
		"strconv.Atoi: parsing \"incorrect\": invalid syntax"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}
