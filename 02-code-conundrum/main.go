package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GameGrab struct {
	Red   int
	Blue  int
	Green int
}

type GameMap map[int][]GameGrab

var conundrumDocument string = "/02-code-conundrum/conundrum.txt"

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	gameIdTotals, powerTotals, err := getGameResult(wd + conundrumDocument)
	if err != nil {
		panic(err)
	}

	fmt.Println(gameIdTotals)
	fmt.Println(powerTotals)
}

func getGameResult(fileName string) (int, int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	games := make(GameMap, 0)
	gameIdTotals := 0
	powerTotals := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		game := scanner.Text()

		gameId, gameGrabs, err := parseGameGrabString(game)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		power, valid := validateGameGrabPossible(gameId, gameGrabs)
		if valid {
			gameIdTotals += gameId
		}
		powerTotals += power

		games[gameId] = gameGrabs
	}

	err = scanner.Err()
	if err != nil {
		return 0, 0, err
	}

	return gameIdTotals, powerTotals, nil
}

func validateGameGrabPossible(gameId int, grabs []GameGrab) (int, bool) {
	totalGrabs := GameGrab{}
	for _, grab := range grabs {
		if grab.Red > totalGrabs.Red {
			totalGrabs.Red = grab.Red
		}
		if grab.Green > totalGrabs.Green {
			totalGrabs.Green = grab.Green
		}
		if grab.Blue > totalGrabs.Blue {
			totalGrabs.Blue = grab.Blue
		}

		// previously added them up not minimum amount
		// totalGrabs.Red += grab.Red
		// totalGrabs.Green += grab.Green
		// totalGrabs.Blue += grab.Blue
	}

	power := totalGrabs.Red * totalGrabs.Green * totalGrabs.Blue

	if totalGrabs.Red > 12 || totalGrabs.Green > 13 || totalGrabs.Blue > 14 {
		return power, false
	}
	return power, true
}

func parseGameGrabString(game string) (int, []GameGrab, error) {
	gameSplit := strings.Split(game, ": ")

	// gameId
	gameId := strings.Split(gameSplit[0], " ")[1]
	number, err := strconv.Atoi(gameId)
	if err != nil {
		return 0, nil, err
	}

	// game grabs
	grabs := strings.Split(gameSplit[1], "; ")
	return number, parseGameGrabs(grabs), nil
}

func parseGameGrabs(gameGrabs []string) []GameGrab {
	grabs := make([]GameGrab, len(gameGrabs))

	for i, grab := range gameGrabs {
		cubes := strings.Split(grab, ", ")
		grabs[i] = parseGrabCubes(cubes)
	}

	return grabs
}

func parseGrabCubes(cubes []string) GameGrab {
	grab := GameGrab{}

	for _, gameCube := range cubes {
		cube := strings.Split(gameCube, " ")

		if len(cube) < 2 {
			fmt.Println("CUBE NO VALUE")
			continue
		}

		number, err := strconv.Atoi(cube[0])
		if err != nil {
			fmt.Println("CUBE NO NUMBER")
			continue
		}

		switch cube[1] {
		case "red":
			grab.Red = number
		case "blue":
			grab.Blue = number
		case "green":
			grab.Green = number
		}
	}

	return grab
}
