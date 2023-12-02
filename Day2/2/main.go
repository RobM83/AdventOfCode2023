package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	id              int
	rounds          []round
	minBalls        balls
	powerOfMinballs int
}

type round struct {
	red   int
	blue  int
	green int
}

type balls struct {
	red   int
	blue  int
	green int
}

func main() {
	gamesInput, _ := readLines("input.txt")
	games := parseInput(gamesInput)
	//games = possibleGames(games)
	//fmt.Println(sumPossibleGameIDs(games))
	games = findMinBalls(games)
	for i := 0; i < len(games); i++ {
		fmt.Println(games[i])
	}
	fmt.Println(sumPowerOfMinBalls(games))
}

func sumPowerOfMinBalls(games []game) int {
	sum := 0
	for i := 0; i < len(games); i++ {
		sum += games[i].powerOfMinballs
	}
	return sum
}

func findMinBalls(games []game) []game {
	for i := 0; i < len(games); i++ {
		for _, rounds := range games[i].rounds {
			if rounds.red > games[i].minBalls.red {
				games[i].minBalls.red = rounds.red
			}
			if rounds.blue > games[i].minBalls.blue {
				games[i].minBalls.blue = rounds.blue
			}
			if rounds.green > games[i].minBalls.green {
				games[i].minBalls.green = rounds.green
			}
		}
		games[i].powerOfMinballs = games[i].minBalls.red * games[i].minBalls.blue * games[i].minBalls.green
	}
	return games
}

func sumPossibleGameIDs(games []game) int {
	sum := 0
	for i := 0; i < len(games); i++ {
		sum += games[i].id
	}
	return sum
}

func possibleGames(games []game) []game {
	maximums := balls{red: 12, blue: 14, green: 13}
	possibleGames := []game{}

	for i := 0; i < len(games); i++ {
		possible := true
		//fmt.Println(games[i])
		for j := 0; j < len(games[i].rounds); j++ {
			//fmt.Println(games[i].rounds[j])
			if games[i].rounds[j].red > maximums.red {
				possible = false
				break
			}
			if games[i].rounds[j].blue > maximums.blue {
				possible = false
				break
			}
			if games[i].rounds[j].green > maximums.green {
				possible = false
				break
			}
		}
		if possible {
			possibleGames = append(possibleGames, games[i])
		}
	}
	return possibleGames
}

func parseInput(input []string) []game {
	var games []game
	for i := 0; i < len(input); i++ {
		games = append(games, parseGame(input[i]))
	}
	return games
}

func parseGame(input string) game {
	//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	gameSplit := strings.Split(input, ":")
	gameID := stringToNumber(strings.Replace(gameSplit[0], "Game ", "", 1))
	rounds := strings.Split(gameSplit[1], ";")

	//Create game
	game := game{id: gameID}
	for i := 0; i < len(rounds); i++ {
		round := parseRound(rounds[i])
		game.rounds = append(game.rounds, round)
	}

	return game
}

func parseRound(input string) round {
	roundSplit := strings.Split(input, ",")
	red := 0
	blue := 0
	green := 0
	for i := 0; i < len(roundSplit); i++ {
		roundSplit[i] = strings.Trim(roundSplit[i], " ")
		roundColorSplit := strings.Split(roundSplit[i], " ")
		color := roundColorSplit[1]
		amount := stringToNumber(roundColorSplit[0])
		switch color {
		case "red":
			red = amount
		case "blue":
			blue = amount
		case "green":
			green = amount
		}
	}
	return round{red: red, blue: blue, green: green}
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
