package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	seeds, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHum, humToLoc := readAndParseInput()
	lowestLoc := math.MaxInt
	for i := 0; i < len(seeds); i = i + 2 {
		start := seeds[i]
		r := seeds[i+1]
		fmt.Println("Processing i: " + strconv.Itoa(i))
		for j := 0; j < r; j++ {
			lowestLoc = getLocationForSeed(j+start, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHum, humToLoc, lowestLoc)
		}
	}

	fmt.Println(lowestLoc)
}

func parseSeedsPart2(line string) []int {
	seeds := strings.Fields(line)[1:]
	var intSeeds = []int{}
	for i := 0; i < len(seeds); i = i + 2 {
		start, _ := strconv.Atoi(seeds[i])
		r, _ := strconv.Atoi(seeds[i+1])
		for j := 0; j < r; j++ {
			intSeeds = append(intSeeds, j+start)
		}
	}
	return intSeeds
}

func getLocationForSeed(seed int, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHum, humToLoc []mapRanges, lowestLoc int) int {
	i := getDestFromMap(seedToSoil, seed)
	i = getDestFromMap(soilToFert, i)
	i = getDestFromMap(fertToWater, i)
	i = getDestFromMap(waterToLight, i)
	i = getDestFromMap(lightToTemp, i)
	i = getDestFromMap(tempToHum, i)
	i = getDestFromMap(humToLoc, i)
	if i < lowestLoc {
		return i
	}
	return lowestLoc
}

func getDestFromMap(mapRanges []mapRanges, index int) int {
	for _, ranges := range mapRanges {
		if ranges.sourceStart <= index && (ranges.sourceStart+ranges.r-1) >= index {
			return ranges.destinationStart + (index - ranges.sourceStart)
		}
	}

	return index
}
