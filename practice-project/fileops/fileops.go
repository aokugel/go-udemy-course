package fileops

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"example.com/practice-project/conversion"
)

type FileManager struct {
	InputValueFilePath string
	OperandValuePath   string
	OutputFilePath     string
}

func New(InputValueFilePath, OperandValuePath, OutputFilePath string) *FileManager {

	return &FileManager{
		InputValueFilePath: InputValueFilePath,
		OperandValuePath:   OperandValuePath,
		OutputFilePath:     OutputFilePath,
	}

}

func (fm *FileManager) GetInputFromFile(mode string) ([]float64, error) {
	var inputMode string
	switch mode {
	case "input":
		inputMode = fm.InputValueFilePath
	case "operand":
		inputMode = fm.OperandValuePath
	default:
		return nil, errors.New("error, please pass a string of either \"input\" or \"operand\" values")
	}
	file, err := os.Open(inputMode)
	if err != nil {
		fmt.Println("error")
		return []float64{}, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	scanner.Err()
	file.Close()
	return conversion.StringsToFloats(lines)
}

// For printing to a straight .txt file

// func OutputCalculations(t *prices.TaxIncludedPriceJob, fileName string) error {
// 	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		f.Close()
// 		return err
// 	}
// 	for k, v := range t.TaxIncludedPrices {
// 		valueText := fmt.Sprintf("Pre-Tax Cost: %v, Tax rate: %%%.2f, Post tax price: %.2f\n", k, t.TaxRate*100, v)
// 		f.Write([]byte(valueText))
// 	}

// 	f.Close()
// 	return nil
// }

func WriteJson(data interface{}, fileName string) error {
	f, err := os.Create(fileName)

	if err != nil {
		f.Close()
		return err
	}
	defer f.Close()
	output_json, err := json.Marshal(data)

	if err != nil {
		return err
	}

	f.Write(output_json)

	return nil
}
