package main

import (
	"testing"
)

func Test_getParameter(test *testing.T) {
	parameter := getParameter("add test", "add")

	if parameter != "test" {
		test.Fail()
	}
}
