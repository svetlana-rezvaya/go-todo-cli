package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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

		if strings.HasPrefix(line, "list") || strings.HasPrefix(line, "find") || strings.HasPrefix(line, "date") {
			filteredNotes, err := filterByMultiCommand(notes, line)
			if err != nil {
				log.Print("unable to filter notes: ", err)
				continue
			}

			text := marshalNotes(filteredNotes)
			fmt.Print(text)
		} else if strings.HasPrefix(line, "add") || strings.HasPrefix(line, "check") || strings.HasPrefix(line, "uncheck") || strings.HasPrefix(line, "delete") {
			updatedNotes, err := updateUsingCommand(notes, line)
			if err != nil {
				log.Print("unable to update notes: ", err)
				continue
			}

			notes = updatedNotes
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
