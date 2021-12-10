package main

import (
	"aoc2021/helper"
	"fmt"
	"os"
	"sort"
)

var costs = map[rune]int {
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var autocosts = map[rune]int {
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func pair(r rune) rune {
	switch r {
	case '[':
		return ']'
	case '(':
		return ')'
	case '<':
		return '>'
	case '{':
		return '}'
	case '}':
		return '{'
	case ')':
		return '('
	case '>':
		return '<'
	case ']':
		return '['
	}
	return ' '
}

func partOne(lines []string) (int, error) {

	var corruptions = map[rune]int{}
	var nest = []rune{}
	for _, line := range lines {
		start:
		for _, ch := range line {
			switch ch {
			case '[', '(', '{', '<':
				nest = append(nest, ch)
			case ']', ')', '}', '>':
				p := pair(ch)
				if nest[len(nest)-1] != p {
					// corrupt
					corruptions[ch]++
					break start
				} else {
					nest = nest[:len(nest)-1]
				}
			}
		}
	}
	score := 0
	for r, c := range corruptions {
		score += c * costs[r]
	}
	return score, nil
}

func partTwo(lines []string) (int, error) {

	var scores = []int{}
	for _, line := range lines {
		nest := []rune{}
		corrupt := false
		start:
		for _, ch := range line {
			switch ch {
			case '[', '(', '{', '<':
				nest = append(nest, ch)
			case ']', ')', '}', '>':
				p := pair(ch)
				if nest[len(nest)-1] != p {
					// corrupt
					corrupt = true
					break start
				} else {
					nest = nest[:len(nest)-1]
				}
			}
		}
		if corrupt {
			continue
		}
		lineScore := 0
		for x := len(nest)-1; x >= 0; x-- {
			p := pair(nest[x])
			lineScore = (lineScore*5) + autocosts[p]

		}
		scores = append(scores, lineScore)
	}
	sort.Ints(scores)
	return scores[(len(scores)/2)], nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := partOne(lines)
	fmt.Printf("Part one: %v\n", ans)
	ans, err = partTwo(lines)
	fmt.Printf("Part two: %v\n", ans)
}
