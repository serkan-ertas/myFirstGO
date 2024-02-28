package note

import (
	"fmt"
)

// used this function to call in main
func NoteFunc() {

	title, content := getNoteData()

	testNote, err := New(title, content)
	if err != nil {
		fmt.Print(err)
		return
	}

	testNote.Display()
	err = testNote.Save()
	if err != nil {
		fmt.Println(err)
		panic(-1)
	}
}
