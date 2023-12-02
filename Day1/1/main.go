package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	calibrationValues, _ := readLines("input.txt")
	correctValues := getCorrectValues(calibrationValues)
	//fmt.Println("Correct values: ", correctValues)
	fmt.Println("Result: ", sumSlice(correctValues))
}

func getCorrectValues(calibrationValues []string) []int {
	result := []int{}
	for _, value := range calibrationValues {
		firstDigit := ""
		lastDigit := ""
		// Loop trough line
		for _, char := range value {
			if unicode.IsDigit(char) {
				// Fill firstValue & lastValue when hitting a digit
				if firstDigit == "" {
					firstDigit = string(char)
					lastDigit = string(char)
				} else {
					// Loop futher and update lastValue when hitting a next digit
					lastDigit = string(char)
				}
			}
		}
		// When done, combine the two numbers
		// Add number to slice
		combinedNumber := firstDigit + lastDigit
		result = append(result, stringToNumber(combinedNumber))
	}

	return result
}

func sumSlice(slice []int) int {
	result := 0
	for _, value := range slice {
		result += value
	}
	return result
}

func stringToNumber(v string) int {
	i, _ := strconv.Atoi(v)
	return i
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
