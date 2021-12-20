package main

import (
	"aoc2021/helper"
	"fmt"
	"math"
	"os"
	"strconv"
)

func algNum(c helper.Coord, data map[helper.Coord]bool) (int, error) {
	neighbours := c.GetOrderedSquare()

	s := ""
	for _, neighbour := range neighbours {
		if v, ok := data[neighbour]; !ok {
			s += "0"
		} else {
			if v {
				s += "1"
			} else {
				s += "0"
			}
		}
	}
	if len(s) != 9 {
		return 0, fmt.Errorf("Bad algorithm conversion")
	}
	num, _ := strconv.ParseInt(s, 2, 64)
	return int(num), nil
}

func getRangeOfPoint(c helper.Coord, data map[helper.Coord]bool) (minx,miny,maxx,maxy int) {
	minx = math.MaxInt32
	miny = math.MaxInt32
	maxx = math.MinInt32
	maxy = math.MinInt32

	mx,my,xmx,xmy := getRanges(data)
	for x := mx; x <= xmx; x++ {
		cc := helper.Coord{X:x, Y:c.Y}
		if data[cc] {
			if cc.X < minx {
				minx = cc.X
			}
			if cc.X > maxx {
				maxx = cc.X
			}
		}
	}
	for y := my; y <= xmy; y++ {
		cc := helper.Coord{X:c.X, Y:y}
		if data[cc] {
			if cc.Y < miny {
				miny = cc.Y
			}
			if cc.Y > maxy {
				maxy = cc.Y
			}
		}
	}
	return minx,miny,maxx,maxy
}

func isEdge(c helper.Coord, data map[helper.Coord]bool) bool {
	mx,my,xmx,xmy := getRanges(data)
	return c.X == mx || c.X == xmx || c.Y == my || c.Y == xmy
}

func problem(algorithm []bool, data map[helper.Coord]bool) (int, error) {

	indata := make(map[helper.Coord]bool)
	for k, v := range data {
		indata[k] = v
	}
	addBorder(&indata)

	for i := 0; i < 2; i++ {
		printMap(indata)
		copymap := make(map[helper.Coord]bool)
		for coord, _ := range indata {
			// get min/max x y ranges for this point
			if isEdge(coord, indata) {
				copymap[coord] = indata[coord]
				continue
			}
			lightnum, err := algNum(coord, indata)
			if err != nil {
				return 0, err
			}
			if lightnum >= len(algorithm) {
				return 0, fmt.Errorf("Bad algorithm lookup")
			}
			copymap[coord] = algorithm[lightnum]
		}

		for k, v := range copymap {
			indata[k] = v
		}

		addBorder(&indata)
	}
	printMap(indata)

	totalLit := 0
	for _, v := range indata {
		if v {
			totalLit++
		}
	}
	return totalLit, nil
}

func readData(lines []string) map[helper.Coord]bool {
	r := make(map[helper.Coord]bool, len(lines))

	for y, line := range lines {
		for x, c := range line {
			coord := helper.Coord{X: x, Y: y}
			if c == '#' {
				r[coord] = true
			} else {
				r[coord] = false
			}
		}
	}
	return r
}

func addBorder(data *map[helper.Coord]bool) {
	mx,my,xmx,xmy := getRanges(*data)
	for y := my-1; y <= xmy+1; y++ {
		for x := mx-1; x <= xmx+1; x++ {
			c := helper.Coord{X: x, Y: y}
			if _, ok := (*data)[c]; !ok {
				(*data)[c] = false
			}
		}
	}
}

func getRanges(data map[helper.Coord]bool) (minx,miny,maxx,maxy int) {
	minx = math.MaxInt32
	miny = math.MaxInt32
	maxx = math.MinInt32
	maxy = math.MinInt32
	for k, _ := range data {
		if k.X < minx {
			minx = k.X
		}
		if k.X > maxx {
			maxx = k.X
		}
		if k.Y < miny {
			miny = k.Y
		}
		if k.Y > maxy {
			maxy = k.Y
		}
	}
	return minx, miny, maxx, maxy
}

func printMap(data map[helper.Coord]bool) {
	mx,my,xmx,xmy := getRanges(data)
	for y := my; y <= xmy; y++ {
		for x := mx; x <= xmx; x++ {
			c := helper.Coord{X:x, Y:y}
			if v, ok := data[c]; ok {
				if v {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")

}

func readAlgorithm(line string) []bool {
	r := make([]bool,len(line))
	for x, c := range line {
		if c == '#' {
			r[x] = true
		} else {
			r[x] = false
		}
	}
	return r
}

func main() {
	fh, _ := os.Open("test.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	algorithm := readAlgorithm(lines[0])
	data := readData(lines[1:])
	ans, err := problem(algorithm, data)
	fmt.Printf("Part one: %v\n", ans)
}
