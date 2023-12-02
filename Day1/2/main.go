package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type writtenValue struct {
	writtenValue string
	digit        string
}

func main() {
	calibrationValues, _ := readLines("input.txt")
	correctValues := getCorrectValues(calibrationValues)
	//fmt.Println("Correct values: ", correctValues)
	fmt.Println("Result: ", sumSlice(correctValues))
}

func getCorrectValues(calibrationValues []string) []int {
	result := []int{}
	for _, value := range calibrationValues {
		fmt.Println("Value: ", value)
		value := rewriteStringValuesByDigits(value)
		fmt.Println("Value: ", value)
		firstDigit := ""
		lastDigit := ""
		for _, char := range value {
			// Loop trough line
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

func rewriteStringValuesByDigits(value string) string {

	writtenValues := []writtenValue{
		{"one", "1"},
		{"two", "2"},
		{"three", "3"},
		{"four", "4"},
		{"five", "5"},
		{"six", "6"},
		{"seven", "7"},
		{"eight", "8"},
		{"nine", "9"},
	}

	value = replaceFromLeftToRight(value, writtenValues)
	value = replaceFromRightToLeft(value, writtenValues)

	return value
}

func replaceFromLeftToRight(value string, writtenValues []writtenValue) string {
	w := ""
	for _, c := range value {
		if unicode.IsDigit(c) {
			return value
		}
		w += string(c)
		for _, writtenValue := range writtenValues {
			if strings.Contains(w, writtenValue.writtenValue) {
				value = strings.Replace(value, writtenValue.writtenValue, writtenValue.digit, 1)
				return value
			}
		}
	}
	return value
}

func replaceFromRightToLeft(value string, writtenValues []writtenValue) string {
	for i := len(value) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(value[i])) {
			return value
		}
		w := value[i:]
		for _, writtenValue := range writtenValues {
			if strings.Contains(w, writtenValue.writtenValue) {
				value = strings.Replace(value, writtenValue.writtenValue, writtenValue.digit, 1)
				break
			}
		}
	}
	return value
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
