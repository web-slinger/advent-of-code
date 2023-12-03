package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var gearRatioSchematic string = "/03-gear-ratios/schematic.txt"

type PotentialPartNumber struct {
	Number     int
	StartIndex int
	EndIndex   int
	LineNumber int
}

// SymbolsMap key = lineIndex_charIndex
type SymbolsMap map[string]string

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	partNumbers, symbols, err := getGearRatios(wd + gearRatioSchematic)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(partNumbers))
	fmt.Println(len(symbols))
}

func getGearRatios(fileName string) ([]PotentialPartNumber, SymbolsMap, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	ppes := make([]PotentialPartNumber, 0)
	symbolsMap := SymbolsMap{}

	scanner := bufio.NewScanner(f)

	index := 0
	for scanner.Scan() {
		line := scanner.Text()

		partNumbers, symbols := getPotentialPartNumbersForLine(line, index)

		ppes = append(ppes, partNumbers...)

		// append line symbols to map
		for key, sym := range symbols {
			symbolsMap[key] = sym
		}
		index++
	}

	err = scanner.Err()
	if err != nil {
		return nil, nil, err
	}

	return ppes, symbolsMap, nil
}

func getPotentialPartNumbersForLine(line string, lineIndex int) ([]PotentialPartNumber, SymbolsMap) {
	ppn := []PotentialPartNumber{}
	symbolMap := SymbolsMap{}

	partNumberBuffer := ""
	for i, char := range line {
		if unicode.IsDigit(char) {
			if i > 0 && partNumberBuffer != "" {
				// check previous char if was digit
				previousWasDigit := line[i-1] == partNumberBuffer[len(partNumberBuffer)-1]

				//reset symbol buffer is previous char wasn't a digit
				if !previousWasDigit {
					partNumberBuffer = ""
				}
			}

			// append char to part number buffer to build up part number
			partNumberBuffer += string(line[i])
		}

		// if char is a 'symbol' then add to map so can be used to validate potential part numbers
		if !unicode.IsDigit(char) && !unicode.IsLetter(char) && string(line[i]) != "." {
			symbolMap[fmt.Sprintf("%d_%d", lineIndex, i)] = string(line[i])
		}

		// add potential part number if current char isn't digit and part buffer isn't empty
		if !unicode.IsDigit(char) && partNumberBuffer != "" {
			number, _ := strconv.Atoi(partNumberBuffer)
			ppn = append(ppn, PotentialPartNumber{
				Number:     number,
				LineNumber: lineIndex,
				StartIndex: i - 1 - (len(partNumberBuffer)),
				EndIndex:   i - 1,
			})
		}

		// reset buffer is current char is not digit
		if !unicode.IsDigit(char) {
			partNumberBuffer = ""
		}
	}
	return ppn, symbolMap
}

func isValidSymbol(potentialPartNumbers []PotentialPartNumber, symbolMap SymbolsMap) int {
	// if line 1 check left, right, down, diag (dr dl)

	// check left right up down all diag (dr dl ur ul)

	// if last line check left, right, up, diag (ur ul)

	// only check 1 index away not multiple diags/ups/downs
	return 0
}
