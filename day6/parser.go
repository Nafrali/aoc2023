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

func readAndParseInput() comps {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	time := scanner.Text()
	scanner.Scan()
	distance := scanner.Text()
	competitions := mapInputToComps(time, distance)

	return competitions
}

func readAndParseInputPart2() comps {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	time := scanner.Text()
	scanner.Scan()
	distance := scanner.Text()
	competitions := mapInputToCompsPart2(time, distance)

	return competitions
}

func mapInputToCompsPart2(time, distance string) comps {
	t, _ := strconv.Atoi(strings.Split(strings.ReplaceAll(time, " ", ""), ":")[1])
	d, _ := strconv.Atoi(strings.Split(strings.ReplaceAll(distance, " ", ""), ":")[1])

	newComp := comp{
		time: t,
		dist: d,
	}
	//competitions := comps{comps: make([]comp, 0)}
	competitions := comps{[]comp{newComp}}

	return competitions
}

func mapInputToComps(time, distance string) comps {
	times := strings.Fields(time)
	distances := strings.Fields(distance)

	competitions := comps{comps: make([]comp, 0)}
	for i := 1; i < len(times); i++ {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distances[i])
		newComp := comp{
			time: t,
			dist: d,
		}
		competitions.comps = append(competitions.comps, newComp)
	}

	return competitions
}
