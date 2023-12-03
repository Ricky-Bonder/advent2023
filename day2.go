package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func day2() {
	lines, err := readLines("day2.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		lineParser(line)
	}

	tot := sumAllGameNumbers(possibleGameNumberSum)
	fmt.Printf("Part One Game Number Total: %v\n", possibleGameNumberSum)
	fmt.Printf("Part One Game Number Total: %v\n", tot)
}

var possibleGameNumberSum []int

func lineParser(line string) {
	fmt.Println("LINE:", line)
	gameNum, AllRoundsInAGame := getGameNumber(line)
	gameNumInt, _ := strconv.Atoi(gameNum)
	possibleGameNumberSum = append(possibleGameNumberSum, gameNumInt)

	separatedRoundsInAGame := getRoundInGame(AllRoundsInAGame)
	var isGamePossible bool = true
	var totalRed int = 0
	var totalBlue int = 0
	var totalGreen int = 0
	for i := 0; i < len(separatedRoundsInAGame); i++ {
		red, blue, green := getCubesCountInASingleRound(separatedRoundsInAGame[i])
		totalRed += red
		totalBlue += blue
		totalGreen += green
		isGamePossible = checkRoundIsPossible(red, blue, green)
		if !isGamePossible && slices.Contains(possibleGameNumberSum, gameNumInt) {
			possibleGameNumberSum = remove(possibleGameNumberSum, gameNumInt)
		}
	}
	if totalRed > 12 && isGamePossible && slices.Contains(possibleGameNumberSum, gameNumInt) {
		possibleGameNumberSum = remove(possibleGameNumberSum, gameNumInt)
	}
	if totalBlue > 14 && isGamePossible && slices.Contains(possibleGameNumberSum, gameNumInt) {
		possibleGameNumberSum = remove(possibleGameNumberSum, gameNumInt)
	}
	if totalGreen > 13 && isGamePossible && slices.Contains(possibleGameNumberSum, gameNumInt) {
		possibleGameNumberSum = remove(possibleGameNumberSum, gameNumInt)
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
			redCubes += count
			fmt.Printf("redCubes: %v\n", redCubes)
		} else if strings.Contains(countAndColor, "blue") {
			countAndColor, _ = strings.CutPrefix(countAndColor, " ")
			countAndColor, _ = strings.CutSuffix(countAndColor, " ")
			singleCountAndColor := strings.Split(countAndColor, " ")
			count, _ := strconv.Atoi(singleCountAndColor[0])
			blueCubes += count
			fmt.Printf("blueCubes: %v\n", blueCubes)
		} else if strings.Contains(countAndColor, "green") {
			countAndColor, _ = strings.CutPrefix(countAndColor, " ")
			countAndColor, _ = strings.CutSuffix(countAndColor, " ")
			singleCountAndColor := strings.Split(countAndColor, " ")
			count, _ := strconv.Atoi(singleCountAndColor[0])
			greenCubes += count
			fmt.Printf("greenCubes: %v\n", greenCubes)
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
