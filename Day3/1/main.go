package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	input, _ := readLines("input.txt")
	engineInput := inputToSlice(input)
	numbers := readEngineSchematic(engineInput)
	fmt.Println("Sum numbers: ", sumNumbers(numbers))
}

func sumNumbers(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}

func readEngineSchematic(engineInput [][]string) []int {
	numbers := make([]int, 0)
	number := ""
	//Read line
	adjecentSymbol := false
	//for _, line := range engineInput {
	for y := 0; y < len(engineInput); y++ {
		if number != "" { //If a new line starts end the previous ended with a number
			if adjecentSymbol {
				//Add number to list
				numbers = append(numbers, stringToNumber(number))
			}
			adjecentSymbol = false
			number = ""
		}
		//Read character
		for x := 0; x < len(engineInput[y]); x++ {
			//If number check for adjacent symbol
			r := []rune(engineInput[y][x])
			if unicode.IsDigit(r[0]) {
				number = number + string(engineInput[y][x])
				//Check left/right/up/down for number
				if !adjecentSymbol {
					adjecentSymbol = checkAdjecentSymbol(engineInput, y, x)
				}
				//Continue to end of number
			} else { //End number
				if adjecentSymbol {
					//Add number to list
					numbers = append(numbers, stringToNumber(number))
				}
				adjecentSymbol = false
				number = ""
				//If adjecent symbol is a number, add to number
			}
		}
	}

	fmt.Println(numbers)
	return numbers
}

func checkAdjecentSymbol(engineInput [][]string, y, x int) bool {
	maxRight := len(engineInput[0]) - 1
	maxDown := len(engineInput) - 1

	//check left
	if x != 0 && isSymbol(engineInput[y][x-1]) {
		return true
	}
	//chedk right
	if x != maxRight && isSymbol(engineInput[y][x+1]) {
		return true
	}
	//check up
	if y != 0 && isSymbol(engineInput[y-1][x]) {
		return true
	}
	//check down
	if y != maxDown && isSymbol(engineInput[y+1][x]) {
		return true
	}
	//check topleft
	if x != 0 && y != 0 && isSymbol(engineInput[y-1][x-1]) {
		return true
	}
	//check topright
	if x != maxRight && y != 0 && isSymbol(engineInput[y-1][x+1]) {
		return true
	}
	//check bottomleft
	if x != 0 && y != maxDown && isSymbol(engineInput[y+1][x-1]) {
		return true
	}
	//check bottomright
	if x != maxRight && y != maxDown && isSymbol(engineInput[y+1][x+1]) {
		return true
	}

	return false
}

func isSymbol(char string) bool {
	r := []rune(char)
	if char == "." || unicode.IsDigit(r[0]) {
		return false
	}
	return true
}

func stringToNumber(v string) int {
	i, _ := strconv.Atoi(v)
	return i
}

func inputToSlice(input []string) [][]string {
	slice := make([][]string, 0)
	for _, line := range input {
		slice = append(slice, stringToSlice(line))
	}
	return slice
}

func stringToSlice(input string) []string {
	slice := make([]string, 0)
	for _, char := range input {
		slice = append(slice, string(char))
	}
	return slice
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
