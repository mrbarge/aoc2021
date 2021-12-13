package main

import (
	"aoc2021/helper"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Fold struct {
	axis string
	pos int
}

func partTwo(coords []helper.Coord, folds []Fold) {

	maxCoord := max(coords)
	grid := make(map[int]map[int]bool, 0)
	for y := 0; y <= maxCoord.Y; y++ {
		for x := 0; x <= maxCoord.X; x++ {
			grid[y] = make(map[int]bool, 0)
			grid[y][x] = false
			if x != 0 && y != 0 {
				grid[-y] = make(map[int]bool, 0)
				grid[-y][-x] = false
				grid[-y][x] = false
				grid[y][-x] = false
			}
		}
	}

	for _, coord := range coords {
		grid[coord.Y][coord.X] = true
	}

	for _, fold := range folds {
		grid = doFold(grid, maxCoord, fold)
	}
	printGrid(grid, maxCoord)
	//grid = doFold(grid, maxCoord, folds[0])
	result := countDots(grid)
	return result, nil
}

func countDots(grid map[int]map[int]bool) (count int) {
	for _, row := range grid {
		for _, v := range row {
			if v {
				count++
			}
		}
	}
	return count
}

func printGrid(grid map[int]map[int]bool, max helper.Coord) {

	minY := math.MaxInt32
	minX := math.MaxInt32
	maxX := math.MinInt32
	maxY := math.MinInt32

	for y := -max.Y; y <= max.Y; y++ {
		for x := -max.X; x <= max.X; x++ {
			if grid[y][x] {
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
			}
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if grid[y][x] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func doFold(grid map[int]map[int]bool, max helper.Coord, f Fold) map[int]map[int]bool {
	if f.axis == "y" {
		// clear dots along the fold line
		for x := -max.X; x <= max.X; x++ {
			grid[f.pos][x] = false
		}
		for y := f.pos+1; y <= max.Y; y++ {
			for x := -max.X; x <= max.X; x++ {
				newY := f.pos - (y - f.pos)
				foldedCoord := helper.Coord{
					X: x,
					Y: newY,
				}
				if grid[y][x] {
					grid[foldedCoord.Y][foldedCoord.X] = grid[y][x]
				}
				grid[y][x] = false
			}
		}
	}
	if f.axis == "x" {
		// clear dots along the fold line
		for y := -max.Y; y <= max.Y; y++ {
			grid[y][f.pos] = false
		}
		for x := f.pos+1; x <= max.X; x++ {
			for y := -max.Y; y <= max.Y; y++ {
				newX := f.pos - (x - f.pos)
				foldedCoord := helper.Coord{
					X: newX,
					Y: y,
				}
				if grid[y][x] {
					grid[foldedCoord.Y][foldedCoord.X] = grid[y][x]
				}
				grid[y][x] = false
			}
		}
	}
	return grid
}

func max(coords []helper.Coord) helper.Coord {
	c := helper.Coord{}
	maxX := 0
	maxY := 0
	for _, coord := range coords {
		if coord.X > maxX {
			maxX = coord.X
		}
		if coord.Y > maxY {
			maxY = coord.Y
		}
	}
	c.X = maxX
	c.Y = maxY
	return c
}

func readData(lines []string) ([]helper.Coord, []Fold) {

	coords := make([]helper.Coord, 0)
	folds := make([]Fold, 0)
	for _, line := range lines {

		if strings.Contains(line, "fold") {
			elems := strings.Split(strings.Split(line, " ")[2], "=")
			n,_ := strconv.Atoi(elems[1])
			f := Fold{
				axis: elems[0],
				pos: n,
			}
			folds = append(folds, f)
		} else {
			elems := strings.Split(line,",")
			x, _ := strconv.Atoi(elems[0])
			y, _ := strconv.Atoi(elems[1])
			coords = append(coords, helper.Coord{X: x, Y: y})
		}
	}
	return coords, folds
}


func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	coords, folds := readData(lines)
	partTwo(coords, folds)
}
