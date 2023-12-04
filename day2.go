package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2() {
	lines, err := readLines("day2.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		// lineParser(line)

		lineParserPart2(line)
	}

	fmt.Printf("Part One Game Number Total: %v\n", possibleGameNumberSum)
	fmt.Printf("Part Two Power sum Total: %v\n", powerSumTogether)
}

var powerSumTogether int

func lineParserPart2(line string) {
	fmt.Println("LINE:", line)
	_, AllRoundsInAGame := getGameNumber(line)

	separatedRoundsInAGame := getRoundInGame(AllRoundsInAGame)
	var maxRed int = 0
	var maxBlue int = 0
	var maxGreen int = 0
	for i := 0; i < len(separatedRoundsInAGame); i++ {
		red, blue, green := getCubesCountInASingleRound(separatedRoundsInAGame[i])
		if red > maxRed {
			maxRed = red
		}
		if blue > maxBlue {
			maxBlue = blue
		}
		if green > maxGreen {
			maxGreen = green
		}
	}
	fmt.Printf("Max Red Cubes: %v\n", maxRed)
	fmt.Printf("Max Green Cubes: %v\n", maxGreen)
	fmt.Printf("Max Blue Cubes: %v\n", maxBlue)
	power := getRoundPower(maxRed, maxBlue, maxGreen)

	powerSumTogether += power
}

func getRoundPower(red int, blue int, green int) int {
	return red * blue * green
}

var possibleGameNumberSum int

func lineParser(line string) {
	fmt.Println("LINE:", line)
	gameNum, AllRoundsInAGame := getGameNumber(line)
	gameNumInt, _ := strconv.Atoi(gameNum)

	separatedRoundsInAGame := getRoundInGame(AllRoundsInAGame)
	var isGamePossible bool = true
	var maxRed int = 0
	var maxBlue int = 0
	var maxGreen int = 0
	for i := 0; i < len(separatedRoundsInAGame); i++ {
		red, blue, green := getCubesCountInASingleRound(separatedRoundsInAGame[i])
		if red > maxRed {
			maxRed = red
		}
		if blue > maxBlue {
			maxBlue = blue
		}
		if green > maxGreen {
			maxGreen = green
		}
	}
	fmt.Printf("Max Red Cubes: %v\n", maxRed)
	fmt.Printf("Max Green Cubes: %v\n", maxGreen)
	fmt.Printf("Max Blue Cubes: %v\n", maxBlue)
	isGamePossible = checkRoundIsPossible(maxRed, maxBlue, maxGreen)
	if isGamePossible {
		possibleGameNumberSum += gameNumInt
	}
}

func sumAllGameNumbers(slice []int) int {
	result := 0
	for _, value := range slice {
		result += value
	}
	return result
}

func remove(slice []int, index int) []int {
	index = index - 1
	if index >= len(slice) {
		return append(slice[:index], 0)
	} else {
		return append(slice[:index], slice[index+1:]...)
	}
}

func checkRoundIsPossible(red int, blue int, green int) bool {
	if red <= 12 && blue <= 14 && green <= 13 {
		return true
	} else {
		return false
	}
}

func getCubesCountInASingleRound(separatedRoundsInAGame string) (int, int, int) {
	var redCubes int = 0
	var blueCubes int = 0
	var greenCubes int = 0

	roundCountsAndColors := strings.Split(separatedRoundsInAGame, ",")
	for _, countAndColor := range roundCountsAndColors {
		if strings.Contains(countAndColor, "red") {
			countAndColor, _ = strings.CutPrefix(countAndColor, " ")
			countAndColor, _ = strings.CutSuffix(countAndColor, " ")
			singleCountAndColor := strings.Split(countAndColor, " ")
			count, _ := strconv.Atoi(singleCountAndColor[0])
			redCubes = count
		} else if strings.Contains(countAndColor, "blue") {
			countAndColor, _ = strings.CutPrefix(countAndColor, " ")
			countAndColor, _ = strings.CutSuffix(countAndColor, " ")
			singleCountAndColor := strings.Split(countAndColor, " ")
			count, _ := strconv.Atoi(singleCountAndColor[0])
			blueCubes = count
		} else if strings.Contains(countAndColor, "green") {
			countAndColor, _ = strings.CutPrefix(countAndColor, " ")
			countAndColor, _ = strings.CutSuffix(countAndColor, " ")
			singleCountAndColor := strings.Split(countAndColor, " ")
			count, _ := strconv.Atoi(singleCountAndColor[0])
			greenCubes = count
		}
	}
	return redCubes, blueCubes, greenCubes
}

func getRoundInGame(line string) []string {
	roundsInGame := strings.Split(line, ";")
	fmt.Printf("Rounds: %v\n", roundsInGame)
	return roundsInGame
}

func getGameNumber(line string) (string, string) {
	gameNumAndCubes := strings.Split(line, ":")
	gameNum := strings.Split(gameNumAndCubes[0], "Game ")
	gameNumber := gameNum[1]
	fmt.Printf("Game Number: %v\n", gameNumber)
	return gameNumber, gameNumAndCubes[1]
}
