package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	storageFilename := flag.String("storage", "storage.data", "storage file")
	flag.Parse()

	notes, err := loadNotes(*storageFilename)
	if err != nil {
		log.Fatal("unable to load notes: ", err)
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
			parameter := getParameter(line, "list")
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
			query := getParameter(line, "find")
			if query == "" {
				log.Print("query missing in 'find' command")
				continue
			}

			filteredNotes := filterByText(notes, query)
			text := marshalNotes(filteredNotes)
			fmt.Print(text)
		} else if strings.HasPrefix(line, "date") {
			parameter := getParameter(line, "date")
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
			text := getParameter(line, "add")
			if text == "" {
				log.Print("text missing in 'add' command")
				continue
			}

			notes = createNote(notes, text)
			err := saveNotes(*storageFilename, notes)
			if err != nil {
				log.Print("unable to save notes: ", err)
			}
		} else if strings.HasPrefix(line, "check") {
			id, err := getIDParameter(line, "check")
			if err != nil {
				log.Print("unable to get the note ID: ", err)
				continue
			}

			updateStatus(notes, id, true)
			err = saveNotes(*storageFilename, notes)
			if err != nil {
				log.Print("unable to save notes: ", err)
			}
		} else if strings.HasPrefix(line, "uncheck") {
			id, err := getIDParameter(line, "uncheck")
			if err != nil {
				log.Print("unable to get the note ID: ", err)
				continue
			}

			updateStatus(notes, id, false)
			err = saveNotes(*storageFilename, notes)
			if err != nil {
				log.Print("unable to save notes: ", err)
			}
		} else if strings.HasPrefix(line, "delete") {
			id, err := getIDParameter(line, "delete")
			if err != nil {
				log.Print("unable to get the note ID: ", err)
				continue
			}

			notes = deleteNote(notes, id)
			err = saveNotes(*storageFilename, notes)
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
