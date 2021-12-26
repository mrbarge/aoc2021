package main

import (
	"aoc2021/helper"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Register map[string]int
type Monad [14]int

var zstate = make(map[string]bool)

var xAdd = []int64 {
	0: 10,
	1:12,
	2:15,
	3:-9,
	4:15,
	5:10,
	6:14,
	7:-5,
	8:14,
	9:-7,
	10:-12,
	11:-10,
	12:-1,
	13:-11,
}

var zDiv = []int64 {
	0:1,
	1:1,
	2:1,
	3:26,
	4:1,
	5:1,
	6:1,
	7:26,
	8:1,
	9:26,
	10:26,
	11:26,
	12:26,
	13:26,
}

var yAdd = []int64 {
	0:15,
	1:8,
	2:2,
	3:6,
	4:13,
	5:4,
	6:1,
	7:9,
	8:5,
	9:13,
	10:9,
	11:6,
	12:2,
	13:2,
}

func runprogram(m Monad) int64 {
	z := int64(0)
	for i, v := range m {
		x := z % 26
		z = z / zDiv[i]
		x += xAdd[i]
		w := int64(v)
		if x != w {
			z = 26*z + w + yAdd[i]
		}
	}
	return z
}

func isLetter(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func add(a string, b string, r *Register) {
	bn := 0
	if !isLetter(b) {
		bn, _ = strconv.Atoi(b)
	} else {
		bn = (*r)[b]
	}
	(*r)[a] += bn
}

func mul(a string, b string, r *Register) {
	bn := 0
	if !isLetter(b) {
		bn, _ = strconv.Atoi(b)
	} else {
		bn = (*r)[b]
	}
	(*r)[a] *= bn
}

func eql(a string, b string, r *Register) {
	bn := 0
	if !isLetter(b) {
		bn, _ = strconv.Atoi(b)
	} else {
		bn = (*r)[b]
	}
	if (*r)[a] == bn {
		(*r)[a] = 1
	} else {
		(*r)[a] = 0
	}
}

func mod(a string, b string, r *Register) error {
	bn := 0
	if !isLetter(b) {
		bn, _ = strconv.Atoi(b)
	} else {
		bn = (*r)[b]
	}
	if bn == 0 {
		return fmt.Errorf("Divide by 0")
	}
	(*r)[a] = (*r)[a] % bn
	return nil
}

func div(a string, b string, r *Register) error {
	bn := 0
	if !isLetter(b) {
		bn, _ = strconv.Atoi(b)
	} else {
		bn = (*r)[b]
	}
	if bn == 0 {
		return fmt.Errorf("Divide by 0")
	}
	(*r)[a] = int(math.Floor((float64)((*r)[a]) / (float64(bn))))
	return nil
}

func program(m Monad, lines []string) (Monad, bool, error) {
	registers := Register{
		"w":0,"x":0,"y":0,"z":0,
	}
	mc := 0

	stepmap := make(map[string]bool, 0)
	for i, line := range lines {

		sl := fmt.Sprintf("%v-%v",i,registers["z"])
		if v, ok := zstate[sl]; ok {
			return m, v, nil
		}
		elems := strings.Split(line, " ")
		cmd := elems[0]
		reg := elems[1]
		switch cmd {
		case "add":
			add(reg,elems[2],&registers)
			break
		case "mod":
			err := mod(reg,elems[2],&registers)
			if err != nil {
				return m, false, err
			}
		case "div":
			err := div(reg,elems[2],&registers)
			if err != nil {
				return m, false, err
			}
			break
		case "mul":
			mul(reg,elems[2],&registers)
			break
		case "eql":
			eql(reg,elems[2],&registers)
			break
		case "inp":
			registers[reg] = m[mc]
			mc++
			if mc == len(m) {
				stepmap[sl] = false
			}
		}
	}
	if registers["z"] == 0 {
		for k, _ := range stepmap {
			zstate[k] = true
		}
	} else {
		for k, _ := range stepmap {
			zstate[k] = false
		}
	}
	if registers["z"] == 0 {
		return m, true, nil
	} else {
		return m, false, nil
	}
}

func partOne(lines []string) int {

	//m := Monad{1,1,8,1,1,1,1,1,1,1,1,1,1,1}
	m := Monad{1,1,8,1,1,9,5,1,3,1,1,4,8,5}

	done := false
	for !done {
		//monad, result, err := program(m,lines)
		result := runprogram(m)
		if result == 0 {
			fmt.Printf("Valid MONAD: %v\n", m)
		}
		// get next monad
		if m[13] == 9 {
			m[13] = 1
			i := 12
			for i >= 0 {
				if m[i] == 9 {
					m[i] = 1
				} else {
					m[i]++
					break
				}
				i--
			}
		} else {
			m[13]++
		}
	}
	return 0
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
}
