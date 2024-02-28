package todo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	// read the whole text
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	// remove new line
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
