package main

import (
	"advent-of-code/helpers"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var scratchcardDocument string = "/04-scratchcards/scratchcards.txt"

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	points, err := getScratchcardPointsFromFile(wd + scratchcardDocument)
	if err != nil {
		panic(err)
	}

	fmt.Println(points)
}

func getScratchcardPointsFromFile(fileName string) (int, error) {
	lines, err := helpers.GetFileLines(fileName)
	if err != nil {
		return 0, err
	}

	return getScratchcardPoints(lines)
}

func getScratchcardPoints(lines []string) (int, error) {

	totalScratchcardPoints := 0

	for _, line := range lines {

		card := strings.Split(line, ": ")
		if len(card) != 2 {
			return 0, fmt.Errorf("expected card format split by ': '")
		}

		//cardId := i + 1 // can use card to get the card number

		play := strings.Split(card[1], " | ")
		if len(play) != 2 {
			return 0, fmt.Errorf("expected play format split by ' | '")
		}

		scratchcard := play[0]
		winningNumbers := play[1]

		//regex scratch card as see if numbers exist in winningNumbers
		re := regexp.MustCompile(`(\d+)`)
		scratchCardNumbers := re.FindAllString(scratchcard, -1)

		var matches int = 0
		var scratchcardPoints = 0
		var ok bool

		for _, match := range scratchCardNumbers {
			ok = strings.Contains(" "+winningNumbers+" ", " "+match+" ")
			if ok {
				scratchcardPoints *= 2
				if matches == 0 {
					scratchcardPoints = 1
				}
				matches++
			}
		}
		totalScratchcardPoints += scratchcardPoints
		fmt.Printf("%s - %d winning numbers, points %d\n", card[0], matches, scratchcardPoints)
	}

	return totalScratchcardPoints, nil
}

/*
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53

split by ": "

slice 0 = lottery card (use regex to get numbers)
			get number of matches in game string

slice 1 = lottery balls (use regex to get numbers)
*/
