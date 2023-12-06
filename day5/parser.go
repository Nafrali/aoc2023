package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type mapRanges struct {
	destinationStart int
	sourceStart      int
	r                int
}

func readAndParseInput() ([]int, []mapRanges, []mapRanges, []mapRanges, []mapRanges, []mapRanges, []mapRanges, []mapRanges) {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	// optionally, resize scanner's capacity for lines over 64K, see next example
	seeds := parseSeeds(scanner.Text())
	ignoreLine(scanner)

	//seedToSoil := readAndParseSection(scanner)
	seedToSoil := readAndParseSection(scanner)
	soilToFert := readAndParseSection(scanner)
	fertToWater := readAndParseSection(scanner)
	waterToLight := readAndParseSection(scanner)
	lightToTemp := readAndParseSection(scanner)
	tempToHum := readAndParseSection(scanner)
	humToLoc := readAndParseSection(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return seeds, seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHum, humToLoc
}

func ignoreLine(scanner *bufio.Scanner) {
	scanner.Scan()
}

func parseSeeds(line string) []int {
	seeds := strings.Fields(line)[1:]
	var intSeeds = []int{}
	for _, t := range seeds {
		i, _ := strconv.Atoi(t)
		intSeeds = append(intSeeds, i)
	}

	return intSeeds
}

func readAndParseSection(scanner *bufio.Scanner) []mapRanges {
	ranges := make([]mapRanges, 0)
	ignoreLine(scanner)
	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Fields(line)
		if len(tmp) == 0 {
			break
		}
		destionationstart, _ := strconv.Atoi(tmp[0])
		sourceStart, _ := strconv.Atoi(tmp[1])
		fieldRange, _ := strconv.Atoi(tmp[2])
		newRange := mapRanges{
			destinationStart: destionationstart,
			sourceStart:      sourceStart,
			r:                fieldRange,
		}
		ranges = append(ranges, newRange)
	}

	return ranges
}
