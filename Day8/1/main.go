package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type direction struct {
	left  string
	right string
}

func main() {
	input, _ := readLines("input.txt")
	instructions, directions := parseInput(input)

	fmt.Println(instructions)
	fmt.Println(directions)

	steps := executeInstructions(instructions, directions)
	fmt.Println(steps)
}

func executeInstructions(instructions string, directions map[string]direction) int {
	//Return steps
	currentPos := "AAA"
	steps := 0

	for currentPos != "ZZZ" {
		for i := 0; i < len(instructions); i++ {
			instruction := string(instructions[i])
			switch instruction {
			case "L":
				currentPos = directions[currentPos].left
			case "R":
				currentPos = directions[currentPos].right
			}
			steps++
			if currentPos == "ZZZ" {
				break
			}
		}
	}

	return steps
}

func parseInput(input []string) (string, map[string]direction) {
	instructions := input[0]
	directions := make(map[string]direction)

	//AAA = (BBB, BBB)
	for i := 2; i < len(input); i++ {
		line := input[i]
		lineSplit := strings.Split(line, "=")
		key := strings.Trim(lineSplit[0], " ")

		lineSplit[1] = strings.Replace(lineSplit[1], "(", "", -1)
		lineSplit[1] = strings.Replace(lineSplit[1], ")", "", -1)
		lineSplit = strings.Split(lineSplit[1], ", ")

		directions[key] = direction{left: strings.Trim(lineSplit[0], " "), right: strings.Trim(lineSplit[1], " ")}
	}

	return instructions, directions
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
