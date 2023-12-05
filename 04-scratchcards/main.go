package main

import (
	"advent-of-code/helpers"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var scratchcardDocument string = "/04-scratchcards/scratchcards.txt"

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	points, numberOfScratchcards, err := getScratchcardPointsFromFile(wd + scratchcardDocument)
	if err != nil {
		panic(err)
	}

	fmt.Println(points)
	fmt.Println(numberOfScratchcards)
}

func getScratchcardPointsFromFile(fileName string) (int, int, error) {
	lines, err := helpers.GetFileLines(fileName)
	if err != nil {
		return 0, 0, err
	}

	return getScratchcardPoints(lines)
}

func getScratchcardPoints(lines []string) (int, int, error) {

	totalScratchcardPoints := 0
	numberOfScratchcards := 0

	scratchMatches := map[int]int{}

	for i := 0; i < len(lines); i++ {

		card := strings.Split(lines[i], ": ")
		if len(card) != 2 {
			return 0, 0, fmt.Errorf("expected card format split by ': '")
		}

		play := strings.Split(card[1], " | ")
		if len(play) != 2 {
			return 0, 0, fmt.Errorf("expected play format split by ' | '")
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

		if matches > 0 {
			cardNumber, _ := strconv.Atoi(strings.Split(card[0], " ")[1])
			scratchMatches[cardNumber] = matches
			// some logic to either push same scratchcard into lines or keep a map of card and no of copies
		}
	}

	scratchCopies := map[int]int{}

	for key, count := range scratchMatches {
		copies, ok := scratchCopies[key]
		if !ok {
			scratchCopies[key] = 0
		}
		copies += 1
		scratchCopies[key] = copies

		for i := 1; i <= count; i++ {
			copies, ok := scratchCopies[key+i]
			if !ok {
				scratchCopies[key+i] = 0
			}
			scratchCopies[key+i] = copies + 1
		}
	}

	fmt.Printf("scratchMatches - %+v\n", scratchMatches)
	fmt.Printf("scratchCopies - %+v\n", scratchCopies)

	return totalScratchcardPoints, numberOfScratchcards, nil
}

/*
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53

split by ": "

slice 0 = lottery card (use regex to get numbers)
			get number of matches in game string

slice 1 = lottery balls (use regex to get numbers)
*/
