package cli

import (
	"errors"
	"strconv"
	"strings"

	"github.com/svetlana-rezvaya/go-todo-cli/encoding"
)

func getParameter(line string, command string) string {
	return strings.TrimSpace(strings.TrimPrefix(line, command))
}

func getIDParameter(line string, command string) (int, error) {
	idStr := getParameter(line, command)
	if idStr == "" {
		return 0, errors.New("note ID missing")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("unable to parse the note ID: " + err.Error())
	}

	return id - encoding.IDOffsetForAlignment, nil
}
