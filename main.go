package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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

		if regexp.MustCompile("^(list|find|date)").MatchString(line) {
			filteredNotes, err := filterByMultiCommand(notes, line)
			if err != nil {
				log.Print("unable to filter notes: ", err)
				continue
			}

			text := marshalNotes(filteredNotes)
			fmt.Print(text)
		} else if regexp.MustCompile("^(add|(un)?check|delete)").MatchString(line) {
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
