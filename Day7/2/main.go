package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards string
	bid   int
	pairs []int
	score int
	rank  int
}

func main() {
	input, _ := readLines("input.txt")
	hands := parseInput(input)
	calculateHands(hands)
	calculateRank(hands)
	fmt.Println(totalWinnings(hands))
}

func totalWinnings(hands []*hand) int {
	total := 0
	for _, h := range hands {
		fmt.Printf("Hand: %s, Bid: %d, Score: %d Rank: %d, Winning: %d\n", h.cards, h.bid, h.score, h.rank, h.bid*h.rank)
		total += h.bid * h.rank
	}
	return total
}

func calculateRank(hands []*hand) {
	sort.Slice(hands, func(i, j int) bool {
		return compare(hands[i], hands[j])
	})

	for i := 0; i < len(hands); i++ {
		hands[i].rank = len(hands) - i
	}
}

func compare(h1, h2 *hand) bool {

	hand1Slice := strings.Split(h1.cards, "")
	hand2Slice := strings.Split(h2.cards, "")

	hand1String := strings.Join(hand1Slice, ",")
	hand2String := strings.Join(hand2Slice, ",")

	hand1String = strings.ReplaceAll(hand1String, "A", "14")
	hand1String = strings.ReplaceAll(hand1String, "K", "13")
	hand1String = strings.ReplaceAll(hand1String, "Q", "12")
	hand1String = strings.ReplaceAll(hand1String, "J", "1")
	hand1String = strings.ReplaceAll(hand1String, "T", "10")

	hand2String = strings.ReplaceAll(hand2String, "A", "14")
	hand2String = strings.ReplaceAll(hand2String, "K", "13")
	hand2String = strings.ReplaceAll(hand2String, "Q", "12")
	hand2String = strings.ReplaceAll(hand2String, "J", "1")
	hand2String = strings.ReplaceAll(hand2String, "T", "10")

	hand1Slice = strings.Split(hand1String, ",")
	hand2Slice = strings.Split(hand2String, ",")

	if h1.score == h2.score { //Same score i.e. FullHouse vs FullHouse
		for k := 0; k < len(hand1Slice); k++ {
			if hand1Slice[k] == hand2Slice[k] { //Same card value i.e. 3 vs 3
				continue
			}
			return stringToNumber(hand1Slice[k]) > stringToNumber(hand2Slice[k])
		}
		return true
	} else {
		return h1.score > h2.score
	}
}

func calculateHands(hands []*hand) {
	for _, h := range hands {
		h.calculateScore()
	}
}

func (h *hand) calculateScore() {
	const (
		FIVEOFAKIND  = 7
		FOUROFAKIND  = 6
		FULLHOUSE    = 5
		THREEOFAKIND = 4
		TWOPAIRS     = 3
		PAIR         = 2
		HIGHCARD     = 1
	)

	//Calculate pairs
	cards := h.cards

	//Check if it contains a J and not only J
	if strings.Contains(cards, "J") && strings.Count(cards, "J") != 5 {
		mapCards := map[string]int{}

		//Create map with card count
		for i := 0; i < len(cards); i++ {
			mapCards[string(cards[i])]++
		}

		//Delete J
		delete(mapCards, "J")

		//Find max
		maxCnt := 0
		card := ""
		for k, v := range mapCards {
			if v > maxCnt {
				maxCnt = v
				card = k
			}
		}

		//Replace J with max card
		cards = strings.ReplaceAll(cards, "J", card)

	}

	for len(cards) > 0 {
		chrCnt := strings.Count(cards, string(cards[0]))
		cards = strings.ReplaceAll(cards, string(cards[0]), "")
		h.pairs[chrCnt-1]++
	}

	//Calculate score
	if h.pairs[4] == 1 {
		//Five of a kind
		h.score = FIVEOFAKIND
		return
	}
	if h.pairs[3] == 1 {
		//Four of a kind
		h.score = FOUROFAKIND
		return
	}
	if h.pairs[2] == 1 && h.pairs[1] == 1 {
		//Full house
		h.score = FULLHOUSE
		return
	}
	if h.pairs[2] == 1 {
		//Three of a kind
		h.score = THREEOFAKIND
		return
	}
	if h.pairs[1] == 2 {
		//Two pairs
		h.score = TWOPAIRS
		return
	}
	if h.pairs[1] == 1 {
		//Pair
		h.score = PAIR
		return
	}
	if h.pairs[0] == 5 {
		//High card
		h.score = HIGHCARD
		return
	}
}

func parseInput(input []string) []*hand {
	var hands []*hand
	//var cards []string

	for _, line := range input {
		in := strings.Split(line, " ")
		hand := hand{in[0], stringToNumber(in[1]), make([]int, 5), 0, 0}
		hands = append(hands, &hand)
	}
	return hands
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
