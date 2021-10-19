package main

import (
	"errors"
	"io/ioutil"
	"os"
)

// LoadNotes ...
func LoadNotes(storageFilename string) ([]Note, error) {
	file, err := os.Open(storageFilename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Note{}, nil
		}
		return nil, errors.New("unable to open a storage file: " + err.Error())
	}
	defer file.Close()

	textBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("unable to read a storage file: " + err.Error())
	}

	notes, err := UnmarshalNotes(string(textBytes))
	if err != nil {
		return nil, errors.New("unable to unmarshal a storage file: " + err.Error())
	}

	return notes, nil
}

// SaveNotes ...
func SaveNotes(storageFilename string, notes []Note) error {
	file, err := os.Create(storageFilename)
	if err != nil {
		return errors.New("unable to create a storage file: " + err.Error())
	}
	defer file.Close()

	text := MarshalNotes(notes)
	_, err = file.WriteString(text)
	if err != nil {
		return errors.New("unable to fill a storage file: " + err.Error())
	}

	return nil
}
