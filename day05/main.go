package main

import (
	"aoc2021/helper"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printIntersects(grid [][]int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 0 {
				fmt.Printf(". ")
			} else {
				fmt.Printf("%v ", grid[y][x])
			}
		}
		fmt.Println("")
	}
}

func min(n ...int) int {
	min := n[0]

	for _, i := range n {
		if min > i {
			min = i
		}
	}
	return min
}

func max(n ...int) int {
	max := n[0]

	for _, i := range n {
		if max < i {
			max = i
		}
	}
	return max
}

func parseLine(l string) (from helper.Coord, to helper.Coord) {

	elems := strings.Split(l, " ")
	fromElems := strings.Split(elems[0],",")
	toElems := strings.Split(elems[2],",")

	fromX, _ := strconv.Atoi(fromElems[0])
	fromY, _ := strconv.Atoi(fromElems[1])
	toX , _ := strconv.Atoi(toElems[0])
	toY, _ := strconv.Atoi(toElems[1])

	from = helper.Coord{
		X: fromX,
		Y: fromY,
	}
	to = helper.Coord {
		X: toX,
		Y: toY,
	}

	return from, to
}

func calculateIntersects(lines []string, diagonals bool, gridX int, gridY int) int {

	intersects := make([][]int, gridX)
	for i, _ := range intersects {
		intersects[i] = make([]int, gridY)
	}

	for _, line := range lines {
		xc, yc := parseLine(line)

		if xc.X != yc.X && xc.Y != yc.Y {
			if !diagonals {
				continue
			}

			// yuck..
			if xc.X < yc.X {
				x := xc.X
				if xc.Y < yc.Y {
					for y := xc.Y; y <= yc.Y; y++ {
						intersects[y][x]++
						x++
					}
				}  else {
					for y := xc.Y; y >= yc.Y; y-- {
						intersects[y][x]++
						x++
					}
				}
			} else {
				x := xc.X
				if xc.Y < yc.Y {
					for y := xc.Y; y <= yc.Y; y++ {
						intersects[y][x]++
						x--
					}
				}  else {
					for y := xc.Y; y >= yc.Y; y-- {
						intersects[y][x]++
						x--
					}
				}
			}
		} else if xc.X == yc.X {
			// horizontal
			for y := min(xc.Y,yc.Y); y <= max(xc.Y,yc.Y); y++ {
				intersects[y][xc.X]++
			}
		} else if xc.Y == yc.Y {
			// vertical
			for x := min(xc.X,yc.X); x <= max(xc.X,yc.X); x++ {
				intersects[xc.Y][x]++
			}
		}
		//break
	}

	count := 0
	for _, r := range intersects {
		for _, v := range r {
			if v >= 2 {
				count++
			}
		}
	}
	return count
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh,false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans := calculateIntersects(lines, false, 1000, 1000)
	fmt.Printf("Part one: %v\n", ans)

	ans = calculateIntersects(lines, true, 1000, 1000)
	fmt.Printf("Part two: %v\n", ans)
}
