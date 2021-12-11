package main

import (
	"aoc2021/helper"
	"fmt"
	"os"
)

type Octopus struct {
	energy int
	flashed bool
}

func partOne(grid [][]*Octopus, steps int) (int, error) {

	flashes := 0
	for step := 0; step < steps; step++ {
		for y, row := range grid {
			for x, o := range row {
				if o.flashed {
					// Don't touch a flashed octopus
					continue
				}
				o.energy += 1
				if o.energy > 9 {
					o.energy = 0
					o.flashed = true
					c := helper.Coord{X:x, Y:y}
					doTheFlash(c, grid)
				}
			}
		}
		// count all that flashed
		flashes += countAndResetFlashes(grid)
	}
	return flashes, nil
}

func partTwo(grid [][]*Octopus) (step int, err error) {

	numOctopus := len(grid) * len(grid[0])
	for step = 1; ; step++ {
		for y, row := range grid {
			for x, o := range row {
				if o.flashed {
					// Don't touch a flashed octopus
					continue
				}
				o.energy += 1
				if o.energy > 9 {
					o.energy = 0
					o.flashed = true
					c := helper.Coord{X:x, Y:y}
					doTheFlash(c, grid)
				}
			}
		}
		// count all that flashed
		flashed := countAndResetFlashes(grid)
		if flashed == numOctopus {
			break
		}
	}
	return step, nil
}

func countAndResetFlashes(grid [][]*Octopus) (flashed int) {
	for _, row := range grid {
		for _, o := range row {
			if o.flashed {
				flashed++
				o.flashed = false
			}
		}
	}
	return flashed
}

func doTheFlash(c helper.Coord, grid [][]*Octopus) {
	neighbours := c.GetNeighboursPos(true)

	for _, n := range neighbours {
		if n.X >= len(grid[0]) || n.Y >= len(grid) {
			// outside grid
			continue
		}
		no := grid[n.Y][n.X]
		if no.flashed {
			continue
		}
		no.energy += 1
		if no.energy > 9 {
			no.energy = 0
			no.flashed = true
			doTheFlash(n, grid)
		}
	}
}

func resetFlash(grid [][]*Octopus) {
	for _, row := range grid {
		for _, o := range row {
			o.flashed = false
		}
	}
}

func printGrid(grid [][]*Octopus) {
	for _, row := range grid {
		for _, o := range row {
			fmt.Printf("%v", o.energy)
			if o.flashed {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func readGrid(lines []string) [][]*Octopus {
	grid := make([][]*Octopus, len(lines))
	for y, line := range lines {
		nums, _ := helper.StrLineToIntArray(line)
		row := make([]*Octopus, len(nums))
		for i, n := range nums {
			o := Octopus{energy: n, flashed: false}
			row[i] = &o
		}
		grid[y] = row
	}
	return grid
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	grid := readGrid(lines)
	ans, err := partOne(grid, 100)
	fmt.Printf("Part one: %v\n", ans)

	grid = readGrid(lines)
	ans, err = partTwo(grid)
	fmt.Printf("Part two: %v\n", ans)
}
