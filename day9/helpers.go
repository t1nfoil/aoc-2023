package main

import (
	"bufio"
	"io/fs"
	"os"
)

func openFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func writeFile(fileName string, stringData []string) error {

	var fileData []byte
	for i := 0; i < len(stringData); i++ {
		b := []byte(stringData[i])
		fileData = append(fileData, b...)
	}

	err := os.WriteFile(fileName, fileData, fs.FileMode(os.O_CREATE))
	if err != nil {
		return err
	}

	return nil
}
