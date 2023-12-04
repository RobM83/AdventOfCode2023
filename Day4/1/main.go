package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type scratchCard struct {
	cardNumber     int
	winningNumbers map[int]bool
	numbers        []int
	score          int
}

func main() {
	input, _ := readLines("input.txt")
	cards := parseInput(input)
	totalScore := getTotalScore(cards)
	fmt.Println(totalScore)
}

func getTotalScore(cards []scratchCard) int {
	totalScore := 0
	for _, card := range cards {
		fmt.Println(card.cardNumber, card.score, card.winningNumbers)
		totalScore += card.score
	}
	return totalScore
}

func parseInput(input []string) []scratchCard {
	cards := make([]scratchCard, 0)
	for _, line := range input {
		scratchCard := newScratchCard(line)
		scratchCard.checkNumbers()
		cards = append(cards, scratchCard)
	}
	return cards
}

func newScratchCard(line string) scratchCard {
	scratchCard := scratchCard{}
	scratchCard.winningNumbers = make(map[int]bool)
	scratchCard.numbers = make([]int, 0)
	scratchCard.score = 0

	card := strings.Split(line, ":")

	cardNumber, _ := strconv.Atoi(strings.Trim(strings.Replace(card[0], "Card", "", 1), " "))
	scratchCard.cardNumber = cardNumber

	allNumbers := strings.Split(card[1], "|")
	winningNumbers := strings.Split(strings.Trim(allNumbers[0], " "), " ")
	numbers := strings.Split(strings.Trim(allNumbers[1], " "), " ")

	for _, nr := range winningNumbers {
		if nr != "" {
			number, _ := strconv.Atoi(strings.Trim(nr, " "))
			scratchCard.winningNumbers[number] = true
		}
	}

	for _, nr := range numbers {
		number, _ := strconv.Atoi(nr)
		scratchCard.numbers = append(scratchCard.numbers, number)
	}

	return scratchCard
}

func (s *scratchCard) checkNumbers() {
	for _, number := range s.numbers {
		if s.winningNumbers[number] {
			if s.score == 0 {
				s.score = 1
			} else {
				s.score = s.score * 2
			}
		}
	}
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
