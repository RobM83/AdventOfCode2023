package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type fromTo struct {
	to  string
	key string
}

type mapping struct {
	ranges []mapRanges
}

type mapRanges struct {
	destinationStart int
	sourceStart      int
	seedRange        int
}

func main() {
	input, _ := readLines("input.txt")
	seeds, seedMap, fromToMap := parseInput(input)
	location := findBestSpots(seeds, seedMap, fromToMap)
	fmt.Println("Best location: ", location)
}

func findBestSpots(seeds []int, seedMap map[string]mapping, fromToMap map[string]fromTo) int {
	lowLocations := make(chan int, len(seeds)/2)

	for i := 0; i < len(seeds)/2; i++ {
		//BRUTE FORCE !!!
		go func(i int) {
			//Startseed, range, startseed, range
			startSeed := seeds[i*2]
			endSeed := startSeed + seeds[i*2+1]

			lowestLocation := math.MaxInt
			for s := startSeed; s < endSeed; s++ {
				location := findBestSpotForSeed(s, seedMap, fromToMap)
				if location <= lowestLocation {
					lowestLocation = location
				}
			}
			log.Printf("Worker %d, Location: %d", i, lowestLocation)
			lowLocations <- lowestLocation
		}(i)
	}

	lowestLocation := math.MaxInt
	for i := 0; i < len(seeds)/2; i++ {
		location := <-lowLocations
		if location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func findBestSpotForSeed(seed int, seedMap map[string]mapping, fromToMap map[string]fromTo) int {
	order := []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity"}
	bestSpot := seed
	//	withinRange := false

	for _, step := range order {
		ranges := seedMap[fromToMap[step].key].ranges
		//Find best spot for seed
		for _, r := range ranges {
			//39 0 15
			if bestSpot >= r.sourceStart && bestSpot <= r.sourceStart+r.seedRange {
				bestSpot = r.destinationStart + (bestSpot - r.sourceStart)
				break
			}
		}
		//fmt.Println(fromToMap[step].key, ":", bestSpot)
	}

	return bestSpot
}

func parseInput(input []string) ([]int, map[string]mapping, map[string]fromTo) {
	var seeds []int
	seedMap := make(map[string]mapping)
	fromToMap := make(map[string]fromTo)
	for i := 0; i < len(input); i++ {

		if strings.Contains(input[i], "seeds:") { //Seed input
			seedsLine := strings.Trim(strings.Replace(input[i], "seeds:", "", -1), " ")
			seedsStr := strings.Split(seedsLine, " ")
			for j := 0; j < len(seedsStr); j++ {
				seeds = append(seeds, stringToNumber(seedsStr[j]))
			}
			continue
		}

		if strings.Contains(input[i], "map:") { //Line is heading for mapping
			header := strings.Trim(strings.Replace(input[i], "map:", "", -1), " ")
			mappingName := strings.Split(header, " ")[0]

			fromToMap[strings.Split(mappingName, "-")[0]] = fromTo{
				to:  strings.Split(header, "-")[2],
				key: mappingName,
			}

			//Read following lines till empty line
			var ranges []mapRanges
			i++ //Go to first range
			for input[i] != "" {
				rangeStr := strings.Split(input[i], " ")
				mapRange := mapRanges{
					destinationStart: stringToNumber(rangeStr[0]),
					sourceStart:      stringToNumber(rangeStr[1]),
					seedRange:        stringToNumber(rangeStr[2]),
				}
				ranges = append(ranges, mapRange)
				i++
				if i == len(input) {
					break
				}
			}
			seedMap[mappingName] = mapping{ranges: ranges}
		}

	}
	return seeds, seedMap, fromToMap
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
