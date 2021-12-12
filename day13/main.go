package main

import (
	"aoc2021/helper"
	"fmt"
	"os"
)

func partOne(lines []string) (int, error) {
	return 0, nil
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
