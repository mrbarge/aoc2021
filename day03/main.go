package main

import (
	"aoc2021/helper"
	"fmt"
	"os"
	"strconv"
)

func partOne(lines []string) (int64, error) {

	gamma := ""
	epsilon := ""

	for x := 0; x < len(lines[0]); x++ {
		mc := mostCommon(lines, x)
		if mc {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	iGamma, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		return -1, err
	}
	iEpsilon, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		return -1, err
	}
	return iGamma * iEpsilon, nil
}

func mostCommon(lines []string, pos int) bool {
	zc := 0
	oc := 0
	for _, v := range lines {
		if v[pos] == '0' {
			zc++
		} else {
			oc++
		}
	}
	if zc > oc {
		return false
	} else if zc < oc {
		return true
	} else {
		return true
	}
}

func leastCommon(lines []string, pos int) bool {
	r := mostCommon(lines, pos)
	return !r
}

func partTwo(lines []string) (int64, error) {

	oxygen := make([]string, len(lines))
	scrub := make([]string, len(lines))
	copy(oxygen, lines)
	copy(scrub, lines)

	pos := 0
	for len(oxygen) > 1 {
		temp := oxygen[:0]
		mc := mostCommon(oxygen, pos)
		for _, v := range oxygen {
			bv, _ := strconv.ParseBool(string(v[pos]))
			if bv == mc {
				temp = append(temp, v)
			}
		}
		oxygen = temp
		pos++
	}

	pos = 0
	for len(scrub) > 1 {
		temp := scrub[:0]
		mc := leastCommon(scrub, pos)
		for _, v := range scrub {
			bv, _ := strconv.ParseBool(string(v[pos]))
			if bv == mc {
				temp = append(temp, v)
			}
		}
		scrub = temp
		pos++
	}

	iOxygen, err := strconv.ParseInt(oxygen[0], 2, 64)
	if err != nil {
		return -1, err
	}
	iScrub, err := strconv.ParseInt(scrub[0], 2, 64)
	if err != nil {
		return -1, err
	}
	return iOxygen * iScrub, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh,true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
	}
	ans, err := partOne(lines)
	if err != nil {
		fmt.Printf("Error in part one: %v", err)
		return
	}
	fmt.Printf("Part one: %v\n",ans)
	ans, err = partTwo(lines)
	if err != nil {
		fmt.Printf("Error in part two: %v", err)
		return
	}
	fmt.Printf("Part two: %v\n",ans)

}
