package main

import "strings"

func getParameter(line string, command string) string {
	return strings.TrimSpace(strings.TrimPrefix(line, command))
}
