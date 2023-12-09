package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readAndParseInput() inputs {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	in := inputs{}
	in.rows = make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)
		values := make([]int, 0)
		for _, v := range split {
			i, _ := strconv.Atoi(v)
			values = append(values, i)
		}
		in.rows = append(in.rows, values)
	}

	return in
}
