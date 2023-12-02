package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	eno   = "eno"
	owt   = "owt"
	eerht = "eerht"
	ruof  = "ruof"
	evif  = "evif"
	xis   = "xis"
	neves = "neves"
	thgie = "thgie"
	enin  = "enin"
)

func main() {
	lines, err := readLines("day1.txt")
	if err != nil {
		panic(err)
	}

	// part1(lines)

	part2(lines)

}

func part2(lines []string) int {
	var partTwoTotal int = 0
	for i := 0; i < len(lines); i++ {
		fmt.Println("LINE:", lines[i])
		var isFirstDigitFound bool = false
		var firstNumber string = ""
		var maybeFirstStringDigit string = ""
		for _, char := range lines[i] {
			if unicode.IsDigit(char) && !isFirstDigitFound {
				isFirstDigitFound = true
				firstNumber = string(char)
				fmt.Printf("firstNumber: %v\n", firstNumber)
				maybeFirstStringDigit = ""
				break
			}
			if unicode.IsLetter(char) && !isFirstDigitFound {

				maybeFirstStringDigit += string(char)

				num, _ := regexToNumber(maybeFirstStringDigit, false)
				if num != "" {
					fmt.Println("First number as string: ", num)
					maybeFirstStringDigit = num
					isFirstDigitFound = true
					firstNumber = ""
					break
				}
			}
		}

		reversedString := reverseString(lines[i])

		var isLastDigitFound bool = false
		var lastNumber string = ""
		var maybeLastStringDigit string = ""
		for _, char := range reversedString {
			if unicode.IsDigit(char) && !isLastDigitFound {
				isLastDigitFound = true
				lastNumber = string(char)
				fmt.Printf("lastNumber: %v\n", lastNumber)
				maybeLastStringDigit = ""
				break
			}
			if unicode.IsLetter(char) && !isLastDigitFound {
				maybeLastStringDigit += string(char)

				num, _ := regexToNumber(maybeLastStringDigit, true)
				if num != "" {
					fmt.Println("Last number as string: ", num)
					maybeLastStringDigit = num
					isLastDigitFound = true
					lastNumber = ""
					break
				}
			}
		}

		var firstNumberAsString string = ""
		var lastNumberAsString string = ""
		if "" != firstNumber {
			firstNumberAsString = firstNumber
		} else if "" != maybeFirstStringDigit {
			firstNumberAsString, _ = letterToStringNumber(maybeFirstStringDigit)
		}
		if "" != lastNumber {
			lastNumberAsString = lastNumber
		} else if "" != maybeLastStringDigit {
			lastNumberAsString, _ = letterToStringNumber(maybeLastStringDigit)
		}
		lineTotal := firstNumberAsString + lastNumberAsString
		lineDigitsAsInt, _ := strconv.Atoi(lineTotal)
		fmt.Printf("line calibration: %v\n", lineDigitsAsInt)
		partTwoTotal += lineDigitsAsInt
	}
	fmt.Printf("Part Two Total: %v\n", partTwoTotal)
	return partTwoTotal
}

func regexToNumber(s string, isReverse bool) (string, error) {
	// Convert the input string to lowercase for case-insensitive comparison
	lowercaseInput := strings.ToLower(s)

	// Map of number words to their corresponding integers
	var numberList []string
	numberList = append(numberList,
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	)

	var numberListReverse []string
	numberListReverse = append(numberList,
		"eno",
		"owt",
		"eerht",
		"ruof",
		"evif",
		"xis",
		"neves",
		"thgie",
		"enin",
	)

	// Iterate over the map to find a match in the input string
	if isReverse {
		for _, word := range numberListReverse {
			if strings.Contains(lowercaseInput, word) {
				return word, nil
			}
		}
	} else {
		for _, word := range numberList {
			if strings.Contains(lowercaseInput, word) {
				return word, nil
			}
		}
	}

	// If no match is found, return an error
	return "", fmt.Errorf("unsupported string: %s", s)
}

func stringToInt(s string) (int, error) {
	// Convert the input string to lowercase for case-insensitive comparison
	lowercaseInput := strings.ToLower(s)

	// Map of number words to their corresponding integers
	numberMapReverse := map[string]int{
		"eno":   1,
		"owt":   2,
		"eerht": 3,
		"ruof":  4,
		"evif":  5,
		"xis":   6,
		"neves": 7,
		"thgie": 8,
		"enin":  9,
	}

	// Iterate over the map to find a match in the input string
	for word, value := range numberMapReverse {
		if strings.Contains(lowercaseInput, word) {
			return value, nil
		}
	}

	// If no match is found, return an error
	return 0, fmt.Errorf("unsupported string: %s", s)
}

func letterToStringNumber(s string) (string, error) {
	switch strings.ToLower(s) {
	case "one":
		return "1", nil
	case "two":
		return "2", nil
	case "three":
		return "3", nil
	case "four":
		return "4", nil
	case "five":
		return "5", nil
	case "six":
		return "6", nil
	case "seven":
		return "7", nil
	case "eight":
		return "8", nil
	case "nine":
		return "9", nil
	case "eno":
		return "1", nil
	case "owt":
		return "2", nil
	case "eerht":
		return "3", nil
	case "ruof":
		return "4", nil
	case "evif":
		return "5", nil
	case "xis":
		return "6", nil
	case "neves":
		return "7", nil
	case "thgie":
		return "8", nil
	case "enin":
		return "9", nil
	default:
		return "0", fmt.Errorf("unsupported string: %s", s)
	}
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
