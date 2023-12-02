package helpers

import (
	"os"
)

func OpenFile(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// calibrationLines := make([]string, 0)

	// scanner := bufio.NewScanner(f)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	calibrationLines = append(calibrationLines, line)
	// }

	// err = scanner.Err()
	// if err != nil {
	// 	return nil, err
	// }
	return f, nil
}
