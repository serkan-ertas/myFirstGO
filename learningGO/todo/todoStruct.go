package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// to create file in correct directory
const directory string = "todo/"

// todo struct
type Todo struct {
	Text string `json:"text"`
}

// todo constructor
func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Text: text,
	}, nil
}

// todo displayer
func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

// todo saver
func (todo Todo) Save() error {
	fileName := directory + "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
