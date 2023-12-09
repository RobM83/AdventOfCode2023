package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type historicLine struct {
	line       []int
	lastValues []int
	nextValue  int
	prevValue  int
}

func main() {
	input, _ := readLines("input.txt")
	history := parseInput(input)
	fmt.Println(totalValue(history))
}

func parseInput(input []string) []historicLine {
	var history []historicLine
	for _, l := range input {
		sv, lv := findNextValues(parseLine(l))
		h := historicLine{
			line:       parseLine(l),
			lastValues: lv,
			nextValue:  findNextValue(lv),
			prevValue:  findPreviousValue(sv),
		}
		history = append(history, h)
	}
	return history
}

func totalValue(history []historicLine) int {
	var total int
	for _, h := range history {
		total += h.prevValue
	}
	return total
}

func findNextValue(lastValues []int) int {
	value := lastValues[len(lastValues)-2] //last one is always 0
	for i := len(lastValues) - 3; i >= 0; i-- {
		value = value + lastValues[i]
	}
	//fmt.Println(value)
	return value
}

func findPreviousValue(startValues []int) int {
	value := startValues[len(startValues)-2] //last one is always 0
	for i := len(startValues) - 3; i >= 0; i-- {
		value = startValues[i] - value
		//fmt.Println(value)
	}
	return value
}

func parseLine(line string) []int {
	var parsedLine []int
	for _, v := range strings.Split(line, " ") {
		parsedLine = append(parsedLine, stringToNumber(strings.Trim(v, " ")))
	}
	return parsedLine
}

func findNextValues(line []int) ([]int, []int) {
	fmt.Println(line)
	lastStepSizes := line
	lastValues := []int{line[len(line)-1]}
	startValues := []int{line[0]}
	//0 3 6 9 12 15
	// 3 3 3 3  3
	//  0 0 0  0
	for !sliceAllZeros(lastStepSizes) {
		stepSizes := []int{}
		for i := 0; i < len(lastStepSizes)-1; i++ {
			//Get every step
			stepSizes = append(stepSizes, lastStepSizes[i+1]-lastStepSizes[i])
		}
		fmt.Println(stepSizes)
		startValues = append(startValues, stepSizes[0])
		lastValues = append(lastValues, stepSizes[len(stepSizes)-1]) //15, 3, 0
		lastStepSizes = stepSizes
	}
	fmt.Println(startValues)
	//fmt.Println(lastValues)
	return startValues, lastValues
}

func sliceAllZeros(slice []int) bool {
	for _, v := range slice {
		if v != 0 {
			return false
		}
	}
	return true
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
