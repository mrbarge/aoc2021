package main

import (
	"aoc2021/helper"
	"fmt"
	"os"
)

func problem(data []int, numDays int) int {

	ages := make(map[int]int, 0)
	for _, n := range data {
		ages[n]++
	}

	days := 0
	for days < numDays {
		nextday := make(map[int]int, 0)
		newfish := 0
		for day, numfish := range ages {
			if day == 0 {
				newfish = numfish
				nextday[6] += numfish
			} else {
				nextday[day-1] += numfish
			}
		}
		nextday[8] = newfish
		for k, v := range nextday {
			ages[k] = v
		}
		ages = nextday
		days++
	}

	sum := 0
	for _, v := range ages {
		sum += v
	}
	return sum
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadCSV(fh)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	nums, _ := helper.StrArrayToInt(lines[0])
	ans := problem(nums, 80)
	fmt.Printf("Part one: %v\n", ans)

	ans = problem(nums, 256)
	fmt.Printf("Part two: %v\n", ans)

}
