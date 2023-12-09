package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func readAndParseInput() navigation {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()
	nodes := mapInputToNodes(scanner)
	nav := navigation{}
	nav.instructions = instructions
	nav.nodes = nodes

	return nav
}

func mapInputToNodes(scanner *bufio.Scanner) map[string]node {
	nm := make(map[string]node, 0)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "=")
		name := strings.TrimSpace(split[0])
		nodeMap := split[1]
		nodeMap = strings.ReplaceAll(nodeMap, "(", "")
		nodeMap = strings.ReplaceAll(nodeMap, ")", "")
		nodesFromMap := strings.Split(nodeMap, ",")
		left := strings.TrimSpace(nodesFromMap[0])
		right := strings.TrimSpace(nodesFromMap[1])
		nn := node{
			right: right,
			left:  left,
		}
		nm[name] = nn
	}
	return nm
}
