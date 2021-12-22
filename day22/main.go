package main

import (
	"aoc2021/helper"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	xf, xt int
	yf, yt int
	zf, zt int
	state bool
}

func max(i int, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}

func min(i int, j int) int {
	if i <= j {
		return i
	} else {
		return j
	}
}

func (c Cube) size() int64 {
	return int64(c.xt-c.xf+1) * int64(c.yt-c.yf+1) * int64(c.zt-c.zf+1)
}

func (c Cube) overlap(t Cube) (bool, *Cube) {
	if c.xf > t.xt || c.xt < t.xf || c.yf > t.yt || c.yt < t.yf || c.zf > t.zt || c.zt < t.zf {
		return false, nil
	}
	r := &Cube{max(c.xf,t.xf), min(c.xt,t.xt), max(c.yf,t.yf), min(c.yt,t.yt), max(c.zf,t.zf), min(c.zt,t.zt),false}
	return true, r
}

func posrange(f int, t int, rf int, rt int) (valid bool, from int, to int) {
	if f < rf && t < rf {
		return false, 0, 0
	}
	if f < rf && t >= rf {
		if t <= rt {
			return true, rf, t
		} else {
			return true, rf, rt
		}
	}
	if f > rt {
		return false, 0, 0
	}
	if f >= rf {
		if t <= rt {
			return true, f, t
		}
		if t > rt {
			return true, f, rt
		}
	}
	return false, 0, 0
}

func readData(lines []string) map[int]map[int]map[int]bool {
	r := make(map[int]map[int]map[int]bool, 0)

	for _, line := range lines {
		power := false
		if strings.Contains(line, "on") {
			power = true
		}
		cubes := strings.Split(strings.Split(line, " ")[1], ",")
		xf, _ := strconv.Atoi(strings.Split(strings.Split(cubes[0], "=")[1], "..")[0])
		xt, _ := strconv.Atoi(strings.Split(strings.Split(cubes[0], "=")[1], "..")[1])
		yf, _ := strconv.Atoi(strings.Split(strings.Split(cubes[1], "=")[1], "..")[0])
		yt, _ := strconv.Atoi(strings.Split(strings.Split(cubes[1], "=")[1], "..")[1])
		zf, _ := strconv.Atoi(strings.Split(strings.Split(cubes[2], "=")[1], "..")[0])
		zt, _ := strconv.Atoi(strings.Split(strings.Split(cubes[2], "=")[1], "..")[1])

		v, sx, ex := posrange(xf, xt, -50, 50)
		if !v {
			continue
		}
		v, sy, ey := posrange(yf, yt, -50, 50)
		if !v {
			continue
		}
		v, sz, ez := posrange(zf, zt, -50, 50)
		if !v {
			continue
		}

		for x := sx; x <= ex; x++ {
			if _, ok := r[x]; !ok {
				r[x] = make(map[int]map[int]bool, 0)
			}
			for y := sy; y <= ey; y++ {
				if _, ok := r[x][y]; !ok {
					r[x][y] = make(map[int]bool, 0)
				}
				for z := sz; z <= ez; z++ {
					r[x][y][z] = power
				}
			}
		}
	}
	return r
}

func readDataPartTwo(lines []string) int64 {
	allcubes := make([]Cube,0)

	for _, line := range lines {
		power := false
		if strings.Contains(line, "on") {
			power = true
		}
		cubes := strings.Split(strings.Split(line, " ")[1], ",")
		xf, _ := strconv.Atoi(strings.Split(strings.Split(cubes[0], "=")[1], "..")[0])
		xt, _ := strconv.Atoi(strings.Split(strings.Split(cubes[0], "=")[1], "..")[1])
		yf, _ := strconv.Atoi(strings.Split(strings.Split(cubes[1], "=")[1], "..")[0])
		yt, _ := strconv.Atoi(strings.Split(strings.Split(cubes[1], "=")[1], "..")[1])
		zf, _ := strconv.Atoi(strings.Split(strings.Split(cubes[2], "=")[1], "..")[0])
		zt, _ := strconv.Atoi(strings.Split(strings.Split(cubes[2], "=")[1], "..")[1])

		c := Cube{xf, xt, yf, yt, zf, zt, power}
		allcubes = append(allcubes, c)
	}


	pastcubes := make([]Cube, 0)
	for _, c := range allcubes {
		compcubes := make([]Cube, 0)
		if c.state {
			compcubes = append(compcubes, c)
		}
		for _, pc := range pastcubes {
			v, oc := c.overlap(pc)
			if v {
				oc.state = !pc.state
				compcubes = append(compcubes, *oc)
			}
		}
		pastcubes = append(pastcubes, compcubes...)
	}

	total := int64(0)
	for _, pc := range pastcubes {
		if pc.state {
			total += pc.size()
		} else {
			total -= pc.size()
		}
	}
	fmt.Printf("%v\n", total)
	return total
}

func partOne(cubes map[int]map[int]map[int]bool) (r int) {
	for _, x := range cubes {
		for _, y := range x {
			for _, z := range y {
				if z {
					r++
				}
			}
		}
	}
	return r
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	cubes := readData(lines)
	ans := partOne(cubes)
	fmt.Printf("Part one: %v\n", ans)
	ans2 := readDataPartTwo(lines)
	fmt.Printf("Part two: %v\n", ans2)

}