package helpers

import (
	"bufio"
	"os"
)

func GetFileLines(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}
