package main

import (
	"errors"
	"os"
)

func saveNotes(storageFilename string, notes []note) error {
	file, err := os.Create(storageFilename)
	if err != nil {
		return errors.New("unable to create a storage file: " + err.Error())
	}
	defer file.Close()

	text := marshalNotes(notes)
	_, err = file.WriteString(text)
	if err != nil {
		return errors.New("unable to fill a storage file: " + err.Error())
	}

	return nil
}
