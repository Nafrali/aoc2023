package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var bitFaceValueMapper = make(map[int32]uint16, 16)

func readAndParseInput() pokerhands {
	hands := make([]pokerhand, 0)

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		originalCards := line[0]
		cards := sortCards(originalCards)
		bid, _ := strconv.Atoi(line[1])
		hand := pokerhand{
			cards:         cards,
			originalOrder: originalCards,
			bid:           bid,
		}
		hands = append(hands, hand)
	}

	return hands
}

func initiateFaceValueMapper() {
	for i, v := range CARDS {
		bitFaceValueMapper[v] = 0x8000 >> i
	}
}

func sortCards(cards string) string {
	c := []rune(cards)
	sort.Slice(c, func(i, j int) bool {
		return strings.Index(CARDS, string(c[i])) < strings.Index(CARDS, string(c[j]))
	})
	return string(c)
}
