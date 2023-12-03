package main

import (
	"bufio"
	"fmt"
	"os"
)

var gearRatioDocument string = "/03-gear-ratios/ratios.txt"

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ratios, err := getGearRatios(wd + gearRatioDocument)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(ratios))
}

func getGearRatios(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ratios := make([]string, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		ratio := scanner.Text()
		ratios = append(ratios, ratio)
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return ratios, nil
}
