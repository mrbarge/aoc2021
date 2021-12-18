package main

import (
	"aoc2021/helper"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Fish struct {
	parent  *Fish
	value   int
	left    *Fish
	right   *Fish
	literal bool
}

func (f *Fish) Print() string {
	ret := ""
	if f.left == nil && f.right == nil {
		ret = fmt.Sprintf("%v", f.value)
	} else {
		ret = fmt.Sprintf("[%s,%s]", f.left.Print(), f.right.Print())
	}
	return ret
}

func partOne(lines []string) (int, error) {
	f := readFish(lines[0])
	for _, line := range lines[1:] {
		newfish := readFish(line)
		nxt := addFish(f, newfish)
		f = nxt
	}
	return f.magnitude(), nil
}

func findMiddle(line string) int {
	depth := 0
	for x, r := range line {
		if r == ',' && depth == 1 {
			return x
		}
		if r == '[' {
			depth++
		}
		if r == ']' {
			depth--
		}
	}
	return -1
}

func (f *Fish) Split() {
	f.left = &Fish{
		value:   int(math.Floor(float64(f.value) / float64(2))),
		literal: true,
		parent:  f,
	}
	f.right = &Fish{
		value:   int(math.Ceil(float64(f.value) / float64(2))),
		literal: true,
		parent:  f,
	}
	f.value = 0
	f.literal = false
}

func (f *Fish) searchSplit() bool {
	if f.literal {
		if f.value > 9 {
			f.Split()
			return true
		} else {
			return false
		}
	} else {
		if f.left.searchSplit() {
			return true
		}
		return f.right.searchSplit()
	}
}

func (f *Fish) searchExplode(depth int) bool {
	if f.literal {
		return false
	}
	if depth > 4 {
		if !f.left.literal {
			return f.left.searchExplode(depth + 1)
		}
		if !f.right.literal {
			return f.right.searchExplode(depth + 1)
		}

		// explode
		lv := f.left.value
		rv := f.right.value
		f.addLeft(lv)
		f.addRight(rv)
		if f.parent.left == f {
			f.parent.left = &Fish{
				parent:  f.parent,
				literal: true,
				value:   0,
			}
		}
		if f.parent.right == f {
			f.parent.right = &Fish{
				parent:  f.parent,
				literal: true,
				value:   0,
			}
		}
		return true
	} else {
		if f.left.searchExplode(depth + 1) {
			return true
		} else if f.right.searchExplode(depth + 1) {
			return true
		}
	}
	return false
}

func (f *Fish) addLeft(v int) bool {
	// first find the first non-right parent
	foundParent := false
	parent := f.parent
	from := f
	for !foundParent {
		if parent != nil {
			if parent.left != from {
				// parent is the first non-left parent
				// now find first number
				//fmt.Printf("Found first non-left: %v\n", parent.Print())
				parent.left.addAnywhereLeft(v)
				foundParent = true
			} else {
				// parent is still on the right
				from = parent
				parent = parent.parent
			}
		} else {
			foundParent = true
		}
	}
	return false
}

func (f *Fish) addRight(v int) bool {
	// first find the first non-right parent
	foundParent := false
	parent := f.parent
	from := f
	for !foundParent {
		if parent != nil {
			if parent.right != from {
				// parent is the first non-right parent
				// now find first number
				parent.right.addAnywhereRight(v)
				foundParent = true
			} else {
				// parent is still on the right
				from = parent
				parent = parent.parent
			}
		} else {
			foundParent = true
		}
	}
	return false
}

func (f *Fish) addAnywhereLeft(v int) bool {
	if f.literal {
		f.value += v
		return true
	}
	if f.right.addAnywhereLeft(v) {
		return true
	}
	if f.left.addAnywhereLeft(v) {
		return true
	}
	return false
}

func (f *Fish) addAnywhereRight(v int) bool {
	if f.literal {
		f.value += v
		return true
	}
	if f.left.addAnywhereRight(v) {
		return true
	}
	if f.right.addAnywhereRight(v) {
		return true
	}
	return false
}

func addFish(f1 *Fish, f2 *Fish) *Fish {
	r := &Fish{
		parent:  nil,
		value:   0,
		left:    f1,
		right:   f2,
		literal: false,
	}
	f1.parent = r
	f2.parent = r
	notdone := true
	for notdone {
		exploded := r.searchExplode(1)
		if exploded {
			continue
		}
		splitted := r.searchSplit()
		if !exploded && !splitted {
			notdone = false
		}
	}
	return r
}

func readFish(line string) *Fish {
	f := &Fish{parent: nil}
	midpoint := findMiddle(line)
	if midpoint < 0 {
		f.value, _ = strconv.Atoi(line)
		f.literal = true
	} else {
		leftfish := line[1:midpoint]
		rightfish := line[midpoint+1 : len(line)-1]
		f.left = readFish(leftfish)
		f.left.parent = f
		f.right = readFish(rightfish)
		f.right.parent = f
	}

	return f
}

func (f *Fish) magnitude() int {
	if f.literal {
		return f.value
	} else {
		return f.left.magnitude()*3 + f.right.magnitude()*2
	}
}

func partTwo(lines []string) int {
	max := 0

	for x := 0; x < len(lines); x++ {
		for y := 0; y < len(lines); y++ {
			if x == y {
				continue
			}
			f1 := readFish(lines[x])
			f2 := readFish(lines[y])
			f3 := addFish(f1, f2)
			mag := f3.magnitude()
			if mag > max {
				max = mag
			}
		}
	}
	return max
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

	ans = partTwo(lines)
	fmt.Printf("Part two: %v\n", ans)

}
