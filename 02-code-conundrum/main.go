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

	games, err := getGames(wd + conundrumDocument)
	if err != nil {
		panic(err)
	}

	for gameId, game := range games {
		fmt.Printf("%d - %+v \n", gameId, game)
	}
}

func getGames(fileName string) (GameMap, error) {

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	games := make(GameMap, 0)

	index := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		index++
		gameGrabsString := strings.Split(scanner.Text(), ": ")[1]

		gameId := index + 1
		gameGrabs := parseGameGrabString(gameGrabsString)

		games[gameId] = gameGrabs
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return games, nil
}

func parseGameGrabString(grabString string) []GameGrab {
	grabs := strings.Split(grabString, "; ")

	return parseGameGrab(grabs)
}

func parseGameGrab(gameGrabs []string) []GameGrab {
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
			grab.Red = number
		case "green":
			grab.Green = number
		default:
			fmt.Println("UNKNOWN CUBE COLOUR - " + cube[1])
		}
	}

	return grab
}
