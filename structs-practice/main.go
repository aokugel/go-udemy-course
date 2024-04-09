package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething(false)
	printSomething("What the fuck man")
	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}

	result := add(1, 2)
	fmt.Println(result)

}
func getTodoData() string {
	return getInput("Todo text: ")

}

func printSomething(value interface{}) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float64: ", floatVal)
		return
	}
	// switch value.(type) {
	// case string:
	// 	fmt.Println("String: ", value)
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case bool:
	// 	fmt.Println("Boolean: ", value)
	// default:
	// 	fmt.Println("Something else: ", value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving note failed with ", err)
		return err
	}
	fmt.Println("Note saved sucessfully.")
	return nil
}

func getNoteData() (string, string) {
	title := getInput("Note title: ")
	content := getInput("Note content: ")
	return title, content
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

func add[T int | float64 | string](a, b T) T {
	return a + b
}
