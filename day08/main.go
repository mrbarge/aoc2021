package main

import (
	"aoc2021/helper"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	segments = map[int]int {
		0:6,
		1:2,
		2:5,
		3:5,
		4:4,
		5:5,
		6:6,
		7:3,
		8:7,
		9:6,
	}
)

func partOne(lines []string) (count int) {

	for _, l := range lines {
		output := strings.Split(l," | ")
		outputCodes := strings.Split(output[1], " ")
		for _, code := range outputCodes {
			cl := len(code)
			for x, v := range segments {
				if x == 1 || x == 4 || x == 7 || x == 8 {
					if cl == v {
						count++
					}
				}
			}
		}
	}
	return count
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func commonWires(a string, b string) (common int) {
	for _, v := range a {
		for _, w := range b {
			if v == w {
				common++
			}
		}
	}
	return common
}

func uncommonWires(a string, b string) (uncommon int)  {
	for _, v := range a {
		found := false
		for _, w := range b {
			if v == w {
				found = true
				break
			}
		}
		if !found {
			uncommon++
		}
	}
	return uncommon
}

func partTwo(lines []string) (count int) {


	for _, l := range lines {
		codeMap := make(map[int]string, 0)
		foundCodes := make(map[string]bool, 0)

		output := strings.Split(l," | ")
		outputCodes := strings.Split(output[1], " ")
		signalCodes := strings.Split(output[0], " ")
		for _, code := range signalCodes {
			cl := len(code)
			sortcode := sortString(code)
			for x, v := range segments {
				if x == 1 || x == 4 || x == 7 || x == 8 {
					if cl == v {
						codeMap[x] = sortcode
						foundCodes[sortcode] = true
						break
					}
				}
			}
		}
		for _, code := range signalCodes {
			cl := len(code)
			sortcode := sortString(code)
			if _, ok := foundCodes[sortcode]; ok {
				continue
			}
			if cl == 6 {
				// 0 6 or 9
				// a 6 will not have the signals of a 1
				if commonWires(codeMap[1], sortcode) == 1 {
					codeMap[6] = sortcode
					foundCodes[sortcode] = true
					continue
				}
			}
		}

		for _, code := range signalCodes {
			cl := len(code)
			sortcode := sortString(code)
			if _, ok := foundCodes[sortcode]; ok {
				continue
			}
			if cl == 5 {
				// 2 3 or 5
				// a 3 will have everything in common with 1
				if commonWires(codeMap[1], sortcode) == 2 && uncommonWires(sortcode, codeMap[1]) == 3 {
					codeMap[3] = sortcode
					foundCodes[sortcode] = true
					continue
				}
				// a 5 will have everything in common with 6
				if commonWires(sortcode, codeMap[6]) == 5 {
					codeMap[5] = sortcode
					foundCodes[sortcode] = true
					continue
				}
				// a 2 will have 4 in common with a 6 and 1 not in common
				if commonWires(sortcode, codeMap[6]) == 4 && uncommonWires(sortcode, codeMap[6]) == 1 {
					codeMap[2] = sortcode
					foundCodes[sortcode] = true
					continue
				}
			}
		}

		for _, code := range signalCodes {
			cl := len(code)
			sortcode := sortString(code)
			if _, ok := foundCodes[sortcode]; ok {
				continue
			}
			if cl == 6 {
				// 0 or 9
				// a 9 will have 5 in common with a 5 and 1 not in common
				if commonWires(sortcode, codeMap[5]) == 5 && uncommonWires(sortcode, codeMap[5]) == 1 {
					codeMap[9] = sortcode
				} else {
					codeMap[0] = sortcode
				}
			}
		}

		sdigit := ""
		for _, code := range outputCodes {
			sortcode := sortString(code)
			for k, v := range codeMap {
				if v == sortcode {
					sdigit += strconv.Itoa(k)
				}
			}
		}
		idigit, _ := strconv.Atoi(sdigit)
		count += idigit
	}

	return count
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans := partOne(lines)
	fmt.Printf("Part one: %v\n", ans)

	ans = partTwo(lines)
	fmt.Printf("Part one: %v\n", ans)

}
