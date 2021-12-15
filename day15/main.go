package main

import (

	"aoc2021/helper"
	"fmt"
	"os"
	"strconv"

	"github.com/albertorestifo/dijkstra"
)

type Element struct {
	c helper.Coord
}


func partOne(lines []string) (int, error) {
	weights := make(map[int]map[int]float64)
	for y, row := range lines {
		for x, v := range row {
			weight, _ := strconv.Atoi(string(v))
			if _, ok := weights[y]; !ok {
				weights[y] = make(map[int]float64, 0)
			}
			weights[y][x] = float64(weight)
		}
	}

	g := dijkstra.Graph{}

	for y, row := range lines {
		for x, _ := range row {
			c := helper.Coord{X: x, Y: y}
			nodeVal := fmt.Sprintf("%v,%v",x,y)
			g[nodeVal] = make(map[string]int, 0)

			neighbours := c.GetNeighboursPos(false)
			for _, n := range neighbours {
				if n.X >= len(lines[0]) || n.Y >= len(lines) {
					continue
				}
				neighVal := fmt.Sprintf("%v,%v",n.X,n.Y)
				g[nodeVal][neighVal] = int(weights[n.Y][n.X])
			}
		}
	}

	end := fmt.Sprintf("%v,%v",len(lines[0])-1,len(lines)-1)
	_, score, err := g.Path("0,0", end)
	if err != nil {
		return 0, err
	}
	return score, nil
}

func megaGrid(lines []string) []string {
	newgrid := make([]string, len(lines)*5)
	biggrid := make([][]int, len(lines)*5)
	for x, _ := range biggrid {
		biggrid[x] = make([]int, len(lines[0])*5)
	}
	linelen := len(lines)
	for y, row := range lines {
		rowlen := len(row)
		for x, r := range row {
			iv, _ := strconv.Atoi(string(r))
			biggrid[y][x] = iv
			for i := 1; i < 5; i++ {
				biggrid[y][(rowlen*i)+x] = (iv+i)
				if biggrid[y][(rowlen*i)+x] > 9 {
					biggrid[y][(rowlen*i)+x] = ((iv+i)%10)+1
				}
			}
		}
	}

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(biggrid[0]); x++ {
			for i := 1; i < 5; i++ {
				biggrid[(linelen*i)+y][x] = biggrid[y][x] + i
				if biggrid[(linelen*i)+y][x] > 9 {
					biggrid[(linelen*i)+y][x] = ((biggrid[y][x] + i) % 10)+1
				}
			}
		}
	}

	for y, row := range biggrid {
		newline := ""
		for _, v := range row {
			newline += strconv.Itoa(v)
		}
		newgrid[y] = newline
	}

	return newgrid
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

	megagrid := megaGrid(lines)
	ans, err = partOne(megagrid)
	fmt.Printf("Part two: %v\n", ans)

}
