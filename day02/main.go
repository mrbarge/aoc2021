package main

import (
	"aoc2021/helper"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partTwo(lines []string) (int, error) {
	position := 0
	depth := 0
	aim := 0

	for _, l := range lines {
		movement := strings.Split(l, " ")[0]
		unitStr := strings.Split(l, " ")[1]
		unit, err := strconv.Atoi(unitStr)
		if err != nil {
			return 0, err
		}

		switch movement {
		case "forward":
			position += unit
			depth += aim * unit
		case "down":
			aim += unit
		case "up":
			aim -= unit
		default:
			return 0, fmt.Errorf("invalid movement: %s", movement)
		}
	}

	return position * depth, nil
}

func partOne(lines []string) (int, error) {
	position := 0
	depth := 0

	for _, l := range lines {
		movement := strings.Split(l, " ")[0]
		unitStr := strings.Split(l, " ")[1]
		unit, err := strconv.Atoi(unitStr)
		if err != nil {
			return 0, err
		}

		switch movement {
		case "forward":
			position += unit
		case "down":
			depth += unit
		case "up":
			depth -= unit
		default:
			return 0, fmt.Errorf("invalid movement: %s", movement)
		}
	}

	return position * depth, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	nums, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
	}
	ans, err := partOne(nums)
	fmt.Printf("Part one: %v\n",ans)
	ans, err = partTwo(nums)
	fmt.Printf("Part two: %v\n",ans)

}

