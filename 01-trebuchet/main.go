package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var calibrationDocument string = "/01-trebuchet/calibration.txt"

func main() {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	calibrationLines, err := getCalibrationLines(wd + calibrationDocument)
	if err != nil {
		panic(err)
	}

	sumOfCalibrationValues, err := getSumOfCalibrationValues(calibrationLines)
	if err != nil {
		panic(err)
	}

	fmt.Println(sumOfCalibrationValues)
}

func getCalibrationLines(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	calibrationLines := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		calibrationLines = append(calibrationLines, line)
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return calibrationLines, nil
}

func getSumOfCalibrationValues(calibrationLines []string) (int, error) {
	sumOfCalibrationValues := 0
	for _, line := range calibrationLines {
		calibrationValue, err := getCalibrationValueFromLine(line)
		if err != nil {
			return 0, err
		}
		sumOfCalibrationValues += calibrationValue
	}
	return sumOfCalibrationValues, nil
}

func getCalibrationValueFromLine(calibrationLine string) (int, error) {
	numbers := getDigitsFromLineIncludingStringNumbers(calibrationLine)

	if len(numbers) != 2 {
		return 0, nil
	}

	return strconv.Atoi(numbers[0] + numbers[1])
}

func getDigitsFromLineIncludingStringNumbers(line string) []string {
	digits := []string{}

	// scan string forwards for stringed numbers else check if number
	for i := 0; i < len(line); i++ {
		stringedNumber, ok := isLinePrefixedWithNumber(line[i:])
		if ok {
			digits = append(digits, stringedNumber)
			continue
		}

		_, err := strconv.Atoi(string(line[i]))
		if err != nil {
			continue
		}

		digits = append(digits, string(line[i]))
	}

	if len(digits) > 0 {
		return []string{digits[0], digits[len(digits)-1]}
	}
	return digits
}

func isLinePrefixedWithNumber(line string) (string, bool) {
	switch {
	case strings.HasPrefix(line, "one"):
		return "1", true
	case strings.HasPrefix(line, "two"):
		return "2", true
	case strings.HasPrefix(line, "three"):
		return "3", true
	case strings.HasPrefix(line, "four"):
		return "4", true
	case strings.HasPrefix(line, "five"):
		return "5", true
	case strings.HasPrefix(line, "six"):
		return "6", true
	case strings.HasPrefix(line, "seven"):
		return "7", true
	case strings.HasPrefix(line, "eight"):
		return "8", true
	case strings.HasPrefix(line, "nine"):
		return "9", true
	default:
		return "0", false
	}
}
