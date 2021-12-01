package main

import (
	"aoc2021/helper"
	"fmt"
	"os"
)

func partOne(nums []int) (increases int) {
	n := nums[0]
	for _, depth := range nums[1:] {
		if depth > n {
			increases++
		}
		n = depth
	}
	return increases
}

func slidingWindow(nums []int, pos int) (size int) {
	if pos+2 >= len(nums) {
		return -1
	}
	return nums[pos] + nums[pos+1] + nums[pos+2]
}

func partTwo(nums []int) (increases int) {
	n := slidingWindow(nums, 0)
	for x := 1; x <= len(nums)-3; x++ {
		depth := slidingWindow(nums, x)
		if depth > n {
			increases++
		}
		n = depth
	}
	return increases
}

func main() {
	fh, _ := os.Open("input.txt")
	nums, err := helper.ReadLinesAsInt(fh)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
	}
	ans := partOne(nums)
	fmt.Printf("Part one: %v\n",ans)
	ans = partTwo(nums)
	fmt.Printf("Part two: %v\n",ans)

}
