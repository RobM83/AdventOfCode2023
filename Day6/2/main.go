package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type race struct {
	time      int
	distance  int
	waysToWin int
}

func main() {
	input, _ := readLines("input.txt")
	races := parseInput(input)
	numberOfWaysToWin(races)
	tww := calculateTotalWaysToWin(races)
	println(tww)
}

func calculateTotalWaysToWin(races []*race) int {
	totalWaysToWin := 1
	for _, r := range races {
		totalWaysToWin = totalWaysToWin * r.waysToWin
	}
	return totalWaysToWin
}

func numberOfWaysToWin(races []*race) {
	for _, r := range races {
		r.calcWaysToWin()
	}
}

func (r *race) calcWaysToWin() {
	waysToWin := 0
	for holdTime := 0; holdTime < r.time; holdTime++ {
		speed := holdTime
		distance := speed * (r.time - holdTime)

		if distance > r.distance {
			waysToWin++
		}
		//fmt.Println(distance)
	}
	r.waysToWin = waysToWin
}

func parseInput(lines []string) []*race {
	timeSlice := strings.Split(lines[0], ":")
	distanceSlice := strings.Split(lines[1], ":")

	var races []*race

	timeStr := strings.ReplaceAll(timeSlice[1], " ", "")
	distanceStr := strings.ReplaceAll(distanceSlice[1], " ", "")
	time := stringToNumber(timeStr)
	distance := stringToNumber(distanceStr)

	r := race{
		time:     time,
		distance: distance,
	}
	races = append(races, &r)

	return races
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
