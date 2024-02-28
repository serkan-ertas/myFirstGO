package todo

import (
	"fmt"

	"whatIsThis.com/intface"
)

func TodoFunc() {
	text := getUserInput("Todo text: ")

	// constructor
	testText, err := New(text)
	if err != nil {
		fmt.Print(err)
		return
	}

	// outputting (display & save)
	err = intface.OutputData(testText)
	if err != nil {
		fmt.Println(err)
		return
	}

	// using interface for saving
	err = intface.SaveData(testText)
	if err != nil {
		fmt.Print(err)
		return
	}
}
