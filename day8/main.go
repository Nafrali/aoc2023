package main

import "fmt"

type navigation struct {
	instructions string
	nodes        map[string]node
}

type node struct {
	right string
	left  string
}

func main() {
	nav := readAndParseInput()

	solvePart1(nav)
	solvePart2(nav)
}

func solvePart1(nav navigation) {
	currNode := "AAA"
	steps := 0
	for currNode != "ZZZ" {
		for _, v := range nav.instructions {
			d := string(v)
			if d == "R" {
				currNode = nav.nodes[currNode].right
			} else {
				currNode = nav.nodes[currNode].left
			}
			steps += 1
			if currNode == "ZZZ" {
				fmt.Printf("Steps taken part1: %d\n", steps)
				return
			}
		}
	}
	fmt.Printf("Steps taken part1: %d\n", steps)
}

func solvePart2part(nav navigation, startNode string) int {
	currNode := startNode
	var steps int = 0
	for currNode[2:] != "Z" {
		for _, v := range nav.instructions {
			d := string(v)
			if d == "R" {
				currNode = nav.nodes[currNode].right
			} else {
				currNode = nav.nodes[currNode].left
			}
			steps += 1
			if currNode[2:] == "Z" {
				return steps
			}
		}
	}
	return steps
}

func goInNode(currNode string, inst string, nodes map[string]node) string {
	if inst == "R" {
		return nodes[currNode].right
	} else {
		return nodes[currNode].left
	}
}

func solvePart2(nav navigation) {
	navNodes := getStartNodesPart2(nav.nodes)
	steps := make([]int, 0)
	for _, sn := range navNodes {
		step := solvePart2part(nav, sn)
		steps = append(steps, step)
	}
	var allSteps int = 1
	a := make(map[int]int, 0)
	for _, v := range steps {
		p := PrimeFactors(int(v))
		for _, i := range p {
			a[i] = i
		}
	}

	for _, v := range a {
		allSteps = allSteps * v
	}
	fmt.Printf("Steps taken part2: %d", allSteps)
}

func getStartNodesPart2(nodeMap map[string]node) []string {
	sn := make([]string, 0)
	for i, _ := range nodeMap {
		ending := i[2:]
		if ending == "A" {
			sn = append(sn, i)
		}
	}

	return sn
}

// Get all prime factors of a given number n
// Taken from https://siongui.github.io/2017/05/09/go-find-all-prime-factors-of-integer-number/
func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}
