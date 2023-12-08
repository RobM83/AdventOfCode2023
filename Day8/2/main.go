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
	instructions, startNodes, directions := parseInput(input)

	fmt.Println(instructions)
	fmt.Println(directions)
	fmt.Println(startNodes)

	steps := executeInstructions(instructions, startNodes, directions)
	fmt.Println(steps)
}

func executeInstructions(instructions string, startNodes []string, directions map[string]direction) int {
	//Return steps
	currentPos := startNodes
	steps := 0

	for !checkEnded(currentPos) {
		for i := 0; i < len(instructions); i++ {
			instruction := string(instructions[i])
			for n := 0; n < len(currentPos); n++ {
				switch instruction {
				case "L":
					currentPos[n] = directions[currentPos[n]].left
				case "R":
					currentPos[n] = directions[currentPos[n]].right
				}
			}
			steps++
			if checkEnded(currentPos) {
				break
			}
		}
	}

	return steps
}

func checkEnded(nodes []string) bool {
	for _, node := range nodes {
		if string(node[2]) != "Z" {
			return false
		}
	}
	return true
}

func parseInput(input []string) (string, []string, map[string]direction) {
	instructions := input[0]
	directions := make(map[string]direction)
	startNodes := []string{}

	//AAA = (BBB, BBB)
	for i := 2; i < len(input); i++ {
		line := input[i]
		lineSplit := strings.Split(line, "=")
		key := strings.Trim(lineSplit[0], " ")

		if string(key[2]) == "A" {
			startNodes = append(startNodes, key)
		}

		lineSplit[1] = strings.Replace(lineSplit[1], "(", "", -1)
		lineSplit[1] = strings.Replace(lineSplit[1], ")", "", -1)
		lineSplit = strings.Split(lineSplit[1], ", ")

		directions[key] = direction{left: strings.Trim(lineSplit[0], " "), right: strings.Trim(lineSplit[1], " ")}
	}

	return instructions, startNodes, directions
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
