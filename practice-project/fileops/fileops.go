package fileops

import (
	"bufio"
	"encoding/json"
	"os"
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

func (fm *FileManager) GetInputFromFile() ([]string, []string, error) {
	priceFile, err := os.Open(fm.InputValueFilePath)
	taxFile, err2 := os.Open(fm.InputValueFilePath)
	if err != nil {
		return []string{}, []string{}, err
	}

	if err2 != nil {
		return []string{}, []string{}, err
	}

	defer priceFile.Close()
	defer taxFile.Close()

	scanner := bufio.NewScanner(priceFile)
	var priceLines []string

	for scanner.Scan() {
		priceLines = append(priceLines, scanner.Text())
	}

	err = scanner.Err()
	priceFile.Close()

	scanner = bufio.NewScanner(taxFile)
	var taxLines []string
	for scanner.Scan() {
		taxLines = append(taxLines, scanner.Text())
	}

	err2 = scanner.Err()

	if err != nil {
		return []string{}, []string{}, err
	}

	if err2 != nil {
		return []string{}, []string{}, err
	}

	return priceLines, taxLines, nil
}

func (fm *FileManager) WriteJson(data interface{}) error {
	f, err := os.Create(fm.OutputFilePath)

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
