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

const helpMessage = "" +
	"filtering commands:\n" +
	"  list [done|to do]  show all notes or notes filtered by specified status\n" +
	"  find TEXT          filter notes by specified text and show them\n" +
	"  date TIMESTAMP     filter notes by specified timestamp (in RFC #822 with numeric zone) and show them\n" +
	"updating commands:\n" +
	"  add TEXT           add a new unchecked note with specified text and current timestamp\n" +
	"  check ID           check the note with specified ID\n" +
	"  uncheck ID         uncheck the note with specified ID\n" +
	"  delete ID          delete the note with specified ID\n" +
	"(to combine filtering commands, use the vertical bar ('|'))\n"

func main() {
	storageFilename :=
		flag.String("storage", "storage.data", "path to a storage file")
	flag.Parse()

	notes, err := LoadNotes(*storageFilename)
	if err != nil {
		log.Fatal("unable to load notes: ", err)
	}

	fmt.Print(helpMessage)

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
			filteredNotes, err := FilterByMultiCommand(notes, line)
			if err != nil {
				log.Print("unable to filter notes: ", err)
				continue
			}

			text := MarshalNotes(filteredNotes)
			fmt.Print(text)
		} else if regexp.MustCompile("^(add|(un)?check|delete)").MatchString(line) {
			updatedNotes, err := UpdateUsingCommand(notes, line)
			if err != nil {
				log.Print("unable to update notes: ", err)
				continue
			}

			notes = updatedNotes
			err = SaveNotes(*storageFilename, notes)
			if err != nil {
				log.Print("unable to save notes: ", err)
			}
		} else if line == "help" {
			fmt.Print(helpMessage)
		} else if line == "exit" {
			os.Exit(0)
		} else {
			log.Print("unknown command: ", line)
		}
	}
}
