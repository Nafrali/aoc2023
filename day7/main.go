package main

import (
	"fmt"
	"sort"
	"strings"
)

const CARDS = "AKQJT98765432"

type pokerhands []pokerhand

type pokerhand struct {
	cards         string
	originalOrder string
	bid           int
}

func main() {
	hands := readAndParseInput()
	sort.Sort(hands)
	var sum int64 = 0
	for i, v := range hands {
		sum = int64((i+1)*v.bid) + sum
	}
	fmt.Println(sum)
}

func (p pokerhands) Len() int {
	return len(p)
}

func (p pokerhands) Less(i, j int) bool {
	handAValue := getHandValue(p[i].cards)
	handBValue := getHandValue(p[j].cards)
	if handAValue > handBValue {
		return true
	} else if handBValue > handAValue {
		return false
	} else if handAValue == handBValue {
		return compareOriginalOrderCards(p[i].originalOrder, p[j].originalOrder)
	}
	return !isAMoreThanB(p[i].cards, p[j].cards)

}

func compareOriginalOrderCards(handA, handB string) bool {
	for i := 0; i < 5; i++ {
		if strings.Index(CARDS, string(handA[i])) > strings.Index(CARDS, string(handB[i])) {
			return true
		} else if strings.Index(CARDS, string(handB[i])) > strings.Index(CARDS, string(handA[i])) {
			return false
		}
	}

	return true
}

func isAMoreThanB(cardsA, cardsB string) bool {
	return getTieValueForCards(cardsA) > getTieValueForCards(cardsB)
}

func getTieValueForCards(cards string) int64 {
	var value int64 = 0
	var valueForShiftAndSub int64 = 14
	cardsSorted := []rune(cards)
	sort.Slice(cardsSorted, func(i, j int) bool {
		return strings.Count(cards, string(cardsSorted[i])) > strings.Count(cards, string(cardsSorted[j]))
	})
	for i, c := range cardsSorted {
		a := int64(strings.Index(CARDS, string(c)))
		var b int64 = 4 * (4 - int64(i))
		value = value | valueForShiftAndSub - a<<b
	}
	return value
}

func getDecimalValuesForCards(cards string) int64 {
	var value int64 = 0

	for i, c := range CARDS {
		count := strings.Count(cards, string(c))
		var countBinary int64 = 0xF >> (4 - count)
		value = value | (countBinary << (4 * (14 - i)))
	}
	return value
}

func isFiveOfAKind(cards string) bool {
	return strings.Count(cards, string(cards[0])) == 5
}

func isFourOfAKind(cards string) bool {
	return getDecimalValuesForCards(cards)%15 == 1
}

func isFullHouse(cards string) bool {
	return getDecimalValuesForCards(cards)%15 == 10
}

func isThreeOfAKind(cards string) bool {
	return getDecimalValuesForCards(cards)%15 == 9
}

func isTwoPairs(cards string) bool {
	return getDecimalValuesForCards(cards)%15 == 7
}

func isPair(cards string) bool {
	return getDecimalValuesForCards(cards)%15 == 6
}

func getHandValue(cards string) int {
	if isFiveOfAKind(cards) {
		return 0
	} else if isFourOfAKind(cards) {
		return 1
	} else if isFullHouse(cards) {
		return 2
	} else if isThreeOfAKind(cards) {
		return 3
	} else if isTwoPairs(cards) {
		return 4
	} else if isPair(cards) {
		return 5
	}

	return 6
}

func (p pokerhands) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
