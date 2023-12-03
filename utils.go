package main

import (
	"bufio"
	"os"
)

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
