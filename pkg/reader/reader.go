package reader

import (
	"bufio"
	"os"
)

// ReadData reads content from a specified file path and returns it as a slice of strings.
// It returns an error if the file cannot be opened or read.
func ReadData(filePath string) ([]string, error) {
	file, err := openFile(filePath)
	if err != nil {
		return nil, err
	}
	defer closeFile(file)

	return readLines(file)
}

func openFile(path string) (*os.File, error) {
	return os.Open(path)
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		panic(err)
	}
}

func readLines(file *os.File) ([]string, error) {
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
