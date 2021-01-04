package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		} else if strings.HasPrefix(line, "find") {
			query := strings.TrimSpace(strings.TrimPrefix(line, "find"))
			if query == "" {
				log.Print("query missing in 'find' command")
				continue
			}

			filteredNotes := filterByText(notes, query)
			text := marshalNotes(filteredNotes)
			fmt.Print(text)
		} else if strings.HasPrefix(line, "date") {
			parameter := strings.TrimSpace(strings.TrimPrefix(line, "date"))
			if parameter == "" {
				log.Print("parameter missing in 'date' command")
				continue
			}

			date, err := time.Parse("02 Jan 06", parameter)
			if err != nil {
				log.Print("unable to parse the 'date' command parameter: ", err)
				continue
			}

			filteredNotes := filterByDate(notes, date)
			text := marshalNotes(filteredNotes)
			fmt.Print(text)
		} else if strings.HasPrefix(line, "add") {
			text := strings.TrimSpace(strings.TrimPrefix(line, "add"))
			if text == "" {
				log.Print("text missing in 'add' command")
				continue
			}

			notes = createNote(notes, text)
			err := saveNotes("storage.data", notes)
			if err != nil {
				log.Print("unable to save notes: ", err)
			}
		} else if strings.HasPrefix(line, "check") {
			idStr := strings.TrimSpace(strings.TrimPrefix(line, "check"))
			if idStr == "" {
				log.Print("note ID missing in 'check' command")
				continue
			}

			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Print("unable to parse the note ID: ", err)
				continue
			}

			updateStatus(notes, id-10000, true)
			err = saveNotes("storage.data", notes)
			if err != nil {
				log.Print("unable to save notes: ", err)
			}
		} else if strings.HasPrefix(line, "uncheck") {
			idStr := strings.TrimSpace(strings.TrimPrefix(line, "uncheck"))
			if idStr == "" {
				log.Print("note ID missing in 'uncheck' command")
				continue
			}

			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Print("unable to parse the note ID: ", err)
				continue
			}

			updateStatus(notes, id-10000, false)
			err = saveNotes("storage.data", notes)
			if err != nil {
				log.Print("unable to save notes: ", err)
			}
		} else if strings.HasPrefix(line, "delete") {
			idStr := strings.TrimSpace(strings.TrimPrefix(line, "delete"))
			if idStr == "" {
				log.Print("note ID missing in 'delete' command")
				continue
			}

			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Print("unable to parse the note ID: ", err)
				continue
			}

			notes = deleteNote(notes, id-10000)
			err = saveNotes("storage.data", notes)
			if err != nil {
				log.Print("unable to save notes: ", err)
			}
		} else if line == "exit" {
			os.Exit(0)
		} else {
			log.Print("unknown command: ", line)
		}
	}
}
