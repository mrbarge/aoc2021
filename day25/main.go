package main

import (
	"aoc2021/helper"
	"fmt"
	"os"
)

func copy(a [][]rune) [][]rune {
	r := make([][]rune, len(a))
	for y, rw := range a {
		r[y] = make([]rune, len(rw))
		for x, v := range rw {
			r[y][x] = v
		}
	}
	return r
}

func partOne(lines []string) (int, error) {

	rows := len(lines)
	cols := len(lines[0])

	grid := make([][]rune, rows)
	for y, r := range lines {
		grid[y] = make([]rune, cols)
		for x, v := range r {
			grid[y][x] = v
		}
	}

	moved := true
	stepcount := 0
	for moved {
		moved = false
		tempmap := copy(grid)
		for y := 0; y < rows; y++ {
			for x := 0; x < cols; x++ {
				if grid[y][x] == '>' {
					nx := (x + 1) % cols
					if grid[y][nx] == '.' {
						tempmap[y][nx] = '>'
						tempmap[y][x] = '.'
						moved = true
					}
				}
			}
		}

		tempmap2 := copy(tempmap)
		for y := 0; y < rows; y++ {
			for x := 0; x < cols; x++ {
				if tempmap[y][x] == 'v' {
					ny := (y + 1) % rows
					if tempmap[ny][x] == '.' {
						tempmap2[ny][x] = 'v'
						tempmap2[y][x] = '.'
						moved = true
					}
				}
			}
		}
		grid = copy(tempmap2)
		stepcount++
	}
	return stepcount, nil
}

func print(grid [][]rune) {
	for _, r := range grid {
		for _, v := range r {
			fmt.Printf("%v",string(v))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
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

}
