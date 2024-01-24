package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title := getInput("Note title: ")
	content := getInput("Note content: ")

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
	}

	userNote.Display()

	WriteContentToFile(userNote.Content, userNote.Title)

}

func getInput(prompt string) string {
	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func WriteContentToFile(value string, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
