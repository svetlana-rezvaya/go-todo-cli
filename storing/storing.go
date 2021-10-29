package storing

import (
	"fmt"
	"io/ioutil"
	"os"

	todo "github.com/svetlana-rezvaya/go-todo-cli"
	"github.com/svetlana-rezvaya/go-todo-cli/encoding"
)

// LoadNotes ...
func LoadNotes(storageFilename string) ([]todo.Note, error) {
	file, err := os.Open(storageFilename)
	if err != nil {
		if os.IsNotExist(err) {
			return []todo.Note{}, nil
		}
		return nil, fmt.Errorf("unable to open a storage file: %s", err)
	}
	defer file.Close()

	textBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read a storage file: %s", err)
	}

	notes, err := encoding.UnmarshalNotes(string(textBytes))
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal a storage file: %s", err)
	}

	return notes, nil
}

// SaveNotes ...
func SaveNotes(storageFilename string, notes []todo.Note) error {
	file, err := os.Create(storageFilename)
	if err != nil {
		return fmt.Errorf("unable to create a storage file: %s", err)
	}
	defer file.Close()

	text := encoding.MarshalNotes(notes)
	_, err = file.WriteString(text)
	if err != nil {
		return fmt.Errorf("unable to fill a storage file: %s", err)
	}

	return nil
}
