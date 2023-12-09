package main

import (
	"advent-of-code/helpers"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	Time                int
	Distance            int
	MinButtonHold       int
	ButtonHoldsWhichWin []int
}

var racesFile string = "06-boat-races/input.txt"

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	races, productOfRaceWinWays, err := getRacesFromFile(wd + racesFile)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(races))
	fmt.Println(productOfRaceWinWays)
}

func getRacesFromFile(fileName string) ([]Race, int, error) {

	lines, err := helpers.GetFileLines(fileName)
	if err != nil {
		return nil, 0, err
	}

	if len(lines) != 2 {
		return nil, 0, fmt.Errorf("expected 2 lines")
	}

	timeNumbers, err := helpers.GetNumbersFromString(lines[0])
	if err != nil {
		return nil, 0, err
	}

	distanceNumbers, err := helpers.GetNumbersFromString(lines[1])
	if err != nil {
		return nil, 0, err
	}

	if len(timeNumbers) != len(distanceNumbers) {
		return nil, 0, fmt.Errorf("rows have differing race counts, time:%d, distance:%d", len(timeNumbers), len(distanceNumbers))
	}

	var races = make([]Race, len(timeNumbers))

	productOfRaceWinWays := 0

	for i := range timeNumbers {
		races[i] = Race{
			Time:     timeNumbers[i],
			Distance: distanceNumbers[i],
		}
		races[i].calcButtonHolds()

		productOfRaceWinWays *= len(races[i].ButtonHoldsWhichWin)
		if i == 0 {
			productOfRaceWinWays = len(races[i].ButtonHoldsWhichWin)
		}
	}

	return races, productOfRaceWinWays, nil
}

func getRacesFromFilePart2(fileName string) (Race, error) {

	lines, err := helpers.GetFileLines(fileName)
	if err != nil {
		return Race{}, err
	}

	if len(lines) != 2 {
		return Race{}, fmt.Errorf("expected 2 lines")
	}

	timeNumbers := helpers.GetStringNumbersFromString(lines[0])

	distanceNumbers := helpers.GetStringNumbersFromString(lines[1])

	time, err := strconv.Atoi(strings.Join(timeNumbers, ""))
	if err != nil {
		return Race{}, err
	}
	distance, err := strconv.Atoi(strings.Join(distanceNumbers, ""))
	if err != nil {
		return Race{}, err
	}

	race := Race{
		Time:     time,
		Distance: distance,
	}
	race.calcButtonHolds()

	return race, nil
}

func (r *Race) calcButtonHolds() {
	r.MinButtonHold = math.MaxInt
	r.ButtonHoldsWhichWin = []int{}
	for i := 1; i <= r.Time; i++ {
		holdTime := i
		timeLeft := r.Time - holdTime

		distancePerTime := timeLeft * holdTime
		if distancePerTime > r.Distance {
			if holdTime < r.MinButtonHold {
				r.MinButtonHold = holdTime
			}
			r.ButtonHoldsWhichWin = append(r.ButtonHoldsWhichWin, holdTime)
		}
	}
}
