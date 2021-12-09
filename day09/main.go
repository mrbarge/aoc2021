package main

import (
	"aoc2021/helper"
	"fmt"
	"os"
	"sort"
)


func partOne(lines []string) (int, error) {

	grid := make([][]int, len(lines))
	for i, line := range lines {
		iline, _ := helper.StrLineToIntArray(line)
		grid[i] = iline
	}

	risk := 0
	for y, row := range grid {
		for x, v := range row {
			if isLowest(x,y,grid) {
				risk += 1 + v
			}
		}
	}

	return risk, nil
}

func partTwo(lines []string) (int, error) {

	grid := make([][]int, len(lines))
	for i, line := range lines {
		iline, _ := helper.StrLineToIntArray(line)
		grid[i] = iline
	}

	basinSizes := make([]int, 0)
	for y, row := range grid {
		for x, _ := range row {
			if isLowest(x,y,grid) {
				basinCoords := getFlowNeighbours(x, y, []helper.Coord{}, grid)
				basinSizes = append(basinSizes, len(basinCoords))
			}
		}
	}
	fmt.Println(basinSizes)
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	return basinSizes[0] * basinSizes[1] * basinSizes[2], nil
}

func seenCoord(x helper.Coord, seen []helper.Coord) bool {
	for _, s := range seen {
		if x.X == s.X && x.Y == s.Y {
			return true
		}
	}
	return false
}

func getFlowNeighbours(x int, y int, seen []helper.Coord, grid[][]int) []helper.Coord {
	c := helper.Coord{X: x, Y: y}
	neighbours := c.GetNeighboursPos(false)

	if seenCoord(c, seen) {
		return seen
	}
	seen = append(seen, c)
	for _, n := range neighbours {
		if n.X >= len(grid[y]) || n.Y >= len(grid) {
			// ignore, off the grid
			continue
		}
		if seenCoord(n, seen) {
			continue
		}
		if grid[n.Y][n.X] == 9 || grid[n.Y][n.X] < grid[y][x] {
			// ignore 9 or lower
			continue
		}
		seen = getFlowNeighbours(n.X, n.Y, seen, grid)
	}
	return seen
}

func isLowest(x int, y int, grid [][]int) bool {
	c := helper.Coord{X: x, Y: y}

	cv := grid[y][x]
	neighbours := c.GetNeighboursPos(false)

	allDifferent := true
	for _, n := range neighbours {
		if n.X >= len(grid[y]) || n.Y >= len(grid) {
			// ignore, off the grid
			continue
		}
		if grid[n.Y][n.X] < cv {
			return false
		}
		if grid[n.Y][n.X] != cv {
			allDifferent = allDifferent && true
		} else {
			allDifferent = false
		}
	}

	return true && allDifferent
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
