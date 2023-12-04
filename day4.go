package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Scratchcard struct {
	gameNumber      int
	winningNumbers  []int
	playingNumbers  []int
	matchingNumbers []int
	powerPoints     int
}

func day4() {
	lines, err := readLines("day4.txt")
	if err != nil {
		panic(err)
	}

	scratchcard := make([]Scratchcard, 0, 1000)

	for index, line := range lines {
		scratchcard = append(scratchcard, Scratchcard{})
		re := regexp.MustCompile("([\\d|]+)")

		values := re.FindAllString(line, -1)

		gameNum, _ := strconv.Atoi(values[0])
		scratchcard[index].gameNumber = gameNum

		isPipeFound := false
		for i := 1; i < len(values); i++ {
			if values[i] != "|" && !isPipeFound {
				winningNum, _ := strconv.Atoi(values[i])
				scratchcard[index].winningNumbers = append(scratchcard[index].winningNumbers, winningNum)
			}
			if values[i] == "|" {
				isPipeFound = true
			}
			if values[i] != "|" && isPipeFound {
				playingNums, _ := strconv.Atoi(values[i])
				scratchcard[index].playingNumbers = append(scratchcard[index].playingNumbers, playingNums)
			}
		}

		for i := 0; i < len(scratchcard[index].playingNumbers); i++ {
			for j := 0; j < len(scratchcard[index].winningNumbers); j++ {
				if scratchcard[index].playingNumbers[i] == scratchcard[index].winningNumbers[j] {
					scratchcard[index].matchingNumbers = append(scratchcard[index].matchingNumbers, scratchcard[index].winningNumbers[j])
				}
			}
		}
		scratchcard[index].powerPoints = int(math.Pow(2, float64(len(scratchcard[index].matchingNumbers))))

		scratchcard[index].printGame()

	}

	total := sumAllPowerNumbers(scratchcard, "powerPoints")
	fmt.Printf("Part One Power Sum Total: %v\n", total)

}

func sumAllPowerNumbers(structure []Scratchcard, powerPoints string) int {
	total := 0

	// Iterate over the slice
	for _, s := range structure {
		// Access the specified field and add its value to the total
		switch powerPoints {
		case "powerPoints":
			total += s.powerPoints
		}
	}

	return total
}

func (sc *Scratchcard) printGame() {
	fmt.Printf("Game Number: %v\n", sc.gameNumber)
	fmt.Printf("Winning numbers: %v\n", sc.winningNumbers)
	fmt.Printf("Playing numbers: %v\n", sc.playingNumbers)
	fmt.Printf("Matching numbers: %v\n", sc.matchingNumbers)
	fmt.Printf("Power of 2: %v\n", sc.powerPoints)
}
