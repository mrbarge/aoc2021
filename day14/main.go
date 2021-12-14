package main

import (
	"aoc2021/helper"
	"fmt"
	"os"
	"sort"
	"strings"
)

func partOne(start string, rules map[string]string) (int, error) {

	polymer := start
	for steps := 0; steps < 10; steps++ {
		tmp := ""
		for x := 0; x < len(polymer)-1; x++ {
			comp := string(polymer[x]) + string(polymer[x+1])
			if _, ok := rules[comp]; ok {
				// a rule matched
				tmp += string(polymer[x]) + rules[comp]
			} else {
				tmp += string(polymer[x])
			}
		}
		tmp += string(polymer[len(polymer)-1])
		polymer = tmp
	}

	elemCount := make(map[rune]int, 0)
	for _, r := range polymer {
		elemCount[r]++
	}
	elemCountList := make([]int, 0)
	for _, v := range elemCount {
		elemCountList = append(elemCountList, v)
	}
	sort.Ints(elemCountList)

	ret := elemCountList[len(elemCountList)-1] - elemCountList[0]
	return ret, nil
}

func makeStartPairs(line string) map[string]int {
	r := make(map[string]int, 0)
	for x := 0; x < len(line)-1; x++ {
		s := line[x:x+2]
		r[s]++
	}
	return r
}

func partOneAlt(start string, rules map[string]string) (int, error) {
	pairCounts := makeStartPairs(start)

	for step := 0; step < 40; step++ {
		tmpCounts := make(map[string]int)
		for pair, count := range pairCounts {
			if _, ok := rules[pair]; ok {
				// A rule exists, translate all pairs
				newPair1 := string(pair[0]) + rules[pair]
				newPair2 := rules[pair] + string(pair[1])
				tmpCounts[newPair1] += count
				tmpCounts[newPair2] += count
			} else {
				tmpCounts[pair] = count
			}
		}
		pairCounts = tmpCounts
	}
	fmt.Println(pairCounts)
	letterFreq := make(map[string]int)
	for pair, count := range pairCounts {
		letterFreq[pair[0:1]] += count
	}
	letterFreq[string(start[len(start)-1])]++
	fmt.Println(letterFreq)
	elemCountList := make([]int, 0)
	for _, v := range letterFreq {
		elemCountList = append(elemCountList, v)
	}
	sort.Ints(elemCountList)
	ret := elemCountList[len(elemCountList)-1] - elemCountList[0]
	return ret, nil
}

func readRules(lines []string) map[string]string {
	r := make(map[string]string, 0)
	for _, line := range lines {
		elems := strings.Split(line, " -> ")
		from := elems[0]
		to := elems[1]
		r[from] = to
	}
	return r
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	start := lines[0]
	rules := readRules(lines[1:])
	ans, err := partOneAlt(start, rules)
	fmt.Printf("Part one: %v\n", ans)
}
