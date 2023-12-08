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
	//Return step
	numberOfSteps := make([]int, len(startNodes))

	for i := 0; i < len(startNodes); i++ {
		//Find number of steps per starting node
		fmt.Println("Start node: ", startNodes[i])
		currentPos := startNodes[i]
		steps := 0
		if string(currentPos[2]) == "Z" {
			break
		}

		for string(currentPos[2]) != "Z" {
			for i := 0; i < len(instructions); i++ {
				instruction := string(instructions[i])
				switch instruction {
				case "L":
					currentPos = directions[currentPos].left
				case "R":
					currentPos = directions[currentPos].right
				}
				fmt.Println(" - ", currentPos)
				steps++
			}
		}

		numberOfSteps[i] = steps
	}

	fmt.Println(numberOfSteps)

	//Find least common multiple of all steps
	lcm := getLeastCommonMultiple(numberOfSteps)

	return lcm
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

func getLeastCommonMultiple(numbers []int) int {
	lcm := numbers[0]
	for i := 0; i < len(numbers); i++ {
		num1 := lcm
		num2 := numbers[i]
		gcd := 1
		for num2 != 0 {
			temp := num2
			num2 = num1 % num2
			num1 = temp
		}
		gcd = num1
		lcm = (lcm * numbers[i]) / gcd
	}

	return lcm
}
