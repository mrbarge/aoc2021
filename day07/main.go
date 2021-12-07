package main

import (
	"aoc2021/helper"
	"fmt"
	"math"
	"os"
)

func problem(nums []int, partOne bool) int {
	minCost := math.MaxInt32
	for i, val := range nums {
		cost := 0
		for j, dest := range nums {
			if i == j {
				continue
			}
			diff := int(math.Abs(float64(val-dest)))
			if partOne {
				cost += diff
			} else {
				cost += diff*(diff+1)/2
			}
		}
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadCSV(fh)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	nums, _ := helper.StrArrayToInt(lines[0])
	ans := problem(nums, true)
	fmt.Printf("Part one: %v\n", ans)

	ans = problem(nums, false)
	fmt.Printf("Part two: %v\n", ans)

}
