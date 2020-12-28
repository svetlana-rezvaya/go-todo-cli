package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	createdAt := time.Date(2020, time.December, 9, 21, 5, 20, 123, time.UTC)
	notes := []note{
		note{ID: 100, CreatedAt: createdAt, IsDone: false, Text: "one"},
		note{ID: 101, CreatedAt: createdAt.Add(time.Hour), IsDone: true, Text: "two"},
		note{
			ID:        102,
			CreatedAt: createdAt.Add(2 * time.Hour),
			IsDone:    false,
			Text:      "three",
		},
		note{
			ID:        103,
			CreatedAt: createdAt.Add(3 * time.Hour),
			IsDone:    true,
			Text:      "four",
		},
		note{
			ID:        104,
			CreatedAt: createdAt.Add(4 * time.Hour),
			IsDone:    false,
			Text:      "five",
		},
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Print("unable to read the line: ", err)
			continue
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "list") {
			filteredNotes := []note{}
			parameter := strings.TrimSpace(strings.TrimPrefix(line, "list"))
			if parameter == "done" {
				filteredNotes = filterByStatus(notes, true)
			} else if parameter == "to do" {
				filteredNotes = filterByStatus(notes, false)
			} else if parameter == "" {
				filteredNotes = notes
			} else {
				log.Print("unknown parameter for 'list' command: ", parameter)
				continue
			}

			text := marshalNotes(filteredNotes)
			fmt.Print(text)
		} else if line == "exit" {
			os.Exit(0)
		} else {
			log.Print("unknown command: ", line)
		}
	}
}
