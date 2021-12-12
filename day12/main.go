package main

import (
	"aoc2021/helper"
	"fmt"
	"os"
	"strings"
)

var caveMap = make(map[string]*Cave)
var sCaveMap = make(map[int]string)

type Cave struct {
	id int
	revisit bool
	neighbours []int
}

func partOne() (int, error) {

	startCave := caveMap["start"]

	seen := make([]int, 0)
	paths := navigate(startCave, seen)

	return len(paths), nil
}

func partTwo() (int, error) {

	startCave := caveMap["start"]

	seen := make([]int, 0)
	paths := navigatePartTwo(startCave, seen)

	return len(paths), nil
}

func navigate(cave *Cave, seen []int) [][]int {

	seen = append(seen, cave.id)
	if sCaveMap[cave.id] == "end" {
		return [][]int{
			seen,
		}
	}

	paths := make([][]int, 0)
	for _, n := range cave.neighbours {

		// Decide if we should visit
		canVisit := true
		// Have we visited this neighbour before?
		for _, seenCave := range seen {
			if seenCave == n {
				// yes, can we revisit?
				if caveMap[sCaveMap[seenCave]].revisit {
					canVisit = true
				} else {
					canVisit = false
				}
				break
			}
		}
		if canVisit {
			nextCave := caveMap[sCaveMap[n]]
			pathsFound := navigate(nextCave, seen)
			for _, p := range pathsFound {
				paths = append(paths, p)
			}
		}
	}
	return paths
}

func navigatePartTwo(cave *Cave, seen []int) [][]int {

	seen = append(seen, cave.id)
	if sCaveMap[cave.id] == "end" {
		return [][]int{
			seen,
		}
	}

	paths := make([][]int, 0)
	for _, n := range cave.neighbours {

		// Ignore start
		if sCaveMap[n] == "start" {
			continue
		}

		// Decide if we should visit
		canVisit := true
		// Have we visited this neighbour before?
		for _, seenCave := range seen {
			if seenCave == n {
				// yes, can we revisit?
				if caveMap[sCaveMap[seenCave]].revisit {
					canVisit = true
				} else {
					// It's a small cave, if we have visited a small
					// cave already, we can't
					if alreadyVisitedSmallCave(seen) {
						canVisit = false
					}
				}
				break
			}
		}
		if canVisit {
			nextCave := caveMap[sCaveMap[n]]
			pathsFound := navigatePartTwo(nextCave, seen)
			for _, p := range pathsFound {
				paths = append(paths, p)
			}
		}
	}
	return paths
}

func alreadyVisitedSmallCave(seen []int) bool {

	lowerCounts := make(map[string]bool)
	for _, cave := range seen {
		caveName := sCaveMap[cave]
		if helper.IsUpper(caveName) || caveName == "start" {
			// ignore
			continue
		}
		if _, ok := lowerCounts[caveName]; ok {
			// already seen this cave
			return true
		} else {
			lowerCounts[caveName] = true
		}
	}
	return false
}

func readCaveMap(lines []string) []*Cave {
	caveId := 0
	caves := make([]*Cave, 0)
	for _, line := range lines {
		elems := strings.Split(line,"-")
		start := elems[0]
		end := elems[1]

		if _, ok := caveMap[start]; !ok {
			revisitable := false
			if helper.IsUpper(start) {
				revisitable = true
			}
			c := Cave{id:caveId, neighbours: []int{}, revisit: revisitable}
			caves = append(caves, &c)
			caveMap[start] = &c
			sCaveMap[c.id] = start
			caveId++
		}
		if _, ok := caveMap[end]; !ok {
			revisitable := false
			if helper.IsUpper(end) {
				revisitable = true
			}
			c := Cave{id:caveId, neighbours: []int{}, revisit: revisitable}
			caves = append(caves, &c)
			caveId++
			caveMap[end] = &c
			sCaveMap[c.id] = end
		}

		caves[caveMap[start].id].neighbours = append(caves[caveMap[start].id].neighbours, caveMap[end].id)
		caves[caveMap[end].id].neighbours = append(caves[caveMap[end].id].neighbours, caveMap[start].id)
	}
	return caves
}

func (c *Cave) Print() {
	fmt.Printf("[%v -> ", c.id)
	for _, n := range c.neighbours {
		fmt.Printf("%v ", n)
	}
	fmt.Printf("]")
	if c.revisit {
		fmt.Printf("* ")
	} else {
		fmt.Printf("  ")
	}
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	readCaveMap(lines)
	ans, err := partOne()
	fmt.Printf("Part one: %v\n", ans)

	ans, err = partTwo()
	fmt.Printf("Part two: %v\n", ans)

}
