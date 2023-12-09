package main

import (
	"fmt"
)

type inputs struct {
	rows [][]int
}

func main() {
	input := readAndParseInput()
	solvePart1(&input)
	solvePart2(&input)
}

func solvePart1(input *inputs) {
	extra := 0
	for _, v := range input.rows {
		extra += calcExtraForRow(v)
	}
	fmt.Printf("Solution part 1: %d\n", extra)
}

func solvePart2(input *inputs) {
	extra := 0
	for _, v := range input.rows {
		reverseRow(&v)
		extra += calcExtraForRow(v)
	}
	fmt.Printf("Solution part 2: %d\n", extra)
}

func reverseRow(in *[]int) {
	for i := 0; i < len(*in)/2; i++ {
		(*in)[i], (*in)[len(*in)-1-i] = (*in)[len(*in)-1-i], (*in)[i]
	}
}

func calcExtraForRow(row []int) int {
	if isAllValuesZero(row) {
		return 0
	}
	newRow := make([]int, 0)
	for i := 0; i < len(row)-1; i++ {
		diff := row[i+1] - row[i]
		newRow = append(newRow, diff)
	}
	return row[len(row)-1] + calcExtraForRow(newRow)
}

func isAllValuesZero(row []int) bool {
	for _, v := range row {
		if v != 0 {
			return false
		}
	}
	return true
}
