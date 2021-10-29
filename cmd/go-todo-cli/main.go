package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/svetlana-rezvaya/go-todo-cli/cli"
	"github.com/svetlana-rezvaya/go-todo-cli/encoding"
	"github.com/svetlana-rezvaya/go-todo-cli/storing"
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
	"misc. commands:\n" +
	"  help               show the help message\n" +
	"  exit               exit\n" +
	"(to combine filtering commands, use the vertical bar ('|'))\n"

func main() {
	storageFilename :=
		flag.String("storage", "storage.data", "path to a storage file")
	flag.Parse()

	notes, err := storing.LoadNotes(*storageFilename)
	if err != nil {
		log.Fatalf("unable to load notes: %s", err)
	}

	fmt.Print(helpMessage)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("unable to read the line: %s", err)
			continue
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if regexp.MustCompile("^(list|find|date)").MatchString(line) {
			filteredNotes, err := cli.FilterByMultiCommand(notes, line)
			if err != nil {
				log.Printf("unable to filter notes: %s", err)
				continue
			}

			text := encoding.MarshalNotes(filteredNotes)
			fmt.Print(text)
		} else if regexp.MustCompile("^(add|(un)?check|delete)").MatchString(line) {
			updatedNotes, err := cli.UpdateUsingCommand(notes, line)
			if err != nil {
				log.Printf("unable to update notes: %s", err)
				continue
			}

			notes = updatedNotes
			err = storing.SaveNotes(*storageFilename, notes)
			if err != nil {
				log.Printf("unable to save notes: %s", err)
			}
		} else if line == "help" {
			fmt.Print(helpMessage)
		} else if line == "exit" {
			os.Exit(0)
		} else {
			log.Printf("unknown command: %s", line)
		}
	}
}
