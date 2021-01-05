package main

import (
	"errors"
	"io/ioutil"
	"os"
)

func loadNotes(storageFilename string) ([]note, error) {
	file, err := os.Open(storageFilename)
	if err != nil {
		return nil, errors.New("unable to open a storage file: " + err.Error())
	}
	defer file.Close()

	textBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("unable to read a storage file: " + err.Error())
	}

	notes, err := unmarshalNotes(string(textBytes))
	if err != nil {
		return nil, errors.New("unable to unmarshal a storage file: " + err.Error())
	}

	return notes, nil
}

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
