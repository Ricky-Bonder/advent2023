package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const (
	one   = "one"
	two   = "two"
	three = "three"
	four  = "four"
	five  = "five"
	six   = "six"
	seven = "seven"
	eight = "eight"
	nine  = "nine"
)

func main() {
	lines, err := readLines("day1.txt")
	if err != nil {
		panic(err)
	}

	part1(lines)

	part2(lines)

}

func part2(lines []string) {

}

func part1(lines []string) int {
	var sumTotal int = 0

	for i := 0; i < len(lines); i++ {
		fmt.Println("LINE:", lines[i])
		var firstNumber string = ""
		for _, char := range lines[i] {
			if unicode.IsDigit(char) {
				if firstNumber == "" {
					firstNumber = string(char)
					fmt.Printf("firstNumber: %v\n", firstNumber)
					break
				}
				// numberAppended = numberAppended + string(char)
				// fmt.Printf("Concatenated: %s\n", numberAppended)
			} else if unicode.IsLetter(char) {
				// fmt.Printf("%c is a letter\n", char)
			} else {
				// fmt.Printf("%c is neither a number nor a letter\n", char)
			}
		}

		reversedString := reverseString(lines[i])

		var lastNumber string = ""
		for _, char := range reversedString {
			if unicode.IsDigit(char) {
				if lastNumber == "" {
					lastNumber = string(char)
					fmt.Printf("lastNumber: %v\n", lastNumber)
					break
				}
				// numberAppended = numberAppended + string(char)
				// fmt.Printf("Concatenated: %s\n", numberAppended)
			} else if unicode.IsLetter(char) {
				// fmt.Printf("%c is a letter\n", char)
			} else {
				// fmt.Printf("%c is neither a number nor a letter\n", char)
			}
		}
		lineTotal := firstNumber + lastNumber
		lineDigitsAsInt, _ := strconv.Atoi(lineTotal)
		fmt.Printf("lineTotal: %v\n", lineDigitsAsInt)
		sumTotal += lineDigitsAsInt
	}
	fmt.Printf("Total: %v\n", sumTotal)
	return sumTotal
}

func reverseString(input string) string {
	// Convert the string to a slice of runes
	runes := []rune(input)

	// Get the length of the slice
	length := len(runes)

	// Reverse the order of the runes in the slice
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the slice of runes back to a string
	reversedString := string(runes)

	return reversedString
}

func appendToRight(originalInt, newInt int) int {
	// Calculate the number of digits in the new integer
	digits := 1
	temp := newInt
	for temp >= 10 {
		temp /= 10
		digits *= 10
	}

	// Append the new integer to the right of the original integer
	result := originalInt*digits + newInt

	return result
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
