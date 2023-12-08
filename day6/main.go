package main

import (
	"fmt"
	"strconv"
)

type comp struct {
	time int
	dist int
}

type comps struct {
	comps []comp
}

func main() {
	competitions := readAndParseInput()
	winsMult := 1
	for _, c := range competitions.comps {
		winsMult = winsMult * calculateWinningOptions(c)
	}
	fmt.Println("Part 1 solution: " + strconv.Itoa(winsMult))

	competitions = readAndParseInputPart2()
	winsPart2 := calculateWinningOptions(competitions.comps[0])
	fmt.Println("Part 2 solution: " + strconv.Itoa(winsPart2))

}

func calculateWinningOptions(competition comp) int {
	nbrOfWins := 0
	for i := 0; i <= competition.time; i++ {
		if calculateDistanceForPressTime(i, competition.time-i) > competition.dist {
			nbrOfWins += 1
		}
	}

	return nbrOfWins
}

func calculateDistanceForPressTime(timePressed, timeLeft int) int {
	return 1 * timePressed * timeLeft
}
