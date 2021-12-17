package main

import (
	"aoc2021/helper"
	"fmt"
)

type BBox struct {
	tl helper.Coord
	tr helper.Coord
	bl helper.Coord
	br helper.Coord
}

func partOne(box BBox) (maxY int, velocities int) {

	maxY = 0
	xv := 0
	yv := 0
	for x := 0; x < 300; x++ {
		for y := -300; y < 300; y++ {
			hit, my := simulate(x,y,box)
			if hit {
				if maxY < my {
					maxY = my
					xv = x
					yv = y
				}
				velocities++
			}
		}
	}
	fmt.Printf("xvel %v yvel %v\n",xv,yv)
	return maxY, velocities

}

func simulate(xv int, yv int, box BBox) (hit bool, maxY int) {

	done := false
	c := helper.Coord{X:xv, Y:yv}
	maxY = yv
	for !done {
		//fmt.Printf("Simulation Coord: %v,%v\n",c.X,c.Y)
		if c.Y > maxY {
			maxY = c.Y
		}
		hit, over := hitTarget(c, box)
		if hit {
			return true, maxY
		} else if over {
			done = true
		}

		if xv > 0 {
			xv--
		}
		yv--
		c.X = c.X + xv
		c.Y = c.Y + yv
	}
	return false, maxY
}

func hitTarget(c helper.Coord, box BBox) (hit bool, overstep bool) {
	if c.X >= box.tl.X && c.X <= box.tr.X && c.Y <= box.tl.Y && c.Y >= box.bl.Y {
		return true, false
	}
	if c.X > box.tr.X || c.Y < box.bl.Y {
		return false, true
	}
	return false, false
}

func main() {
	box := BBox{
		tl: helper.Coord{X: 175, Y: -79},
		tr: helper.Coord{X: 227, Y: -79},
		bl: helper.Coord{X: 175, Y: -134},
		br: helper.Coord{X: 227, Y: -134},
	}
	//testbox := BBox{
	//	tl: helper.Coord{X: 20, Y: -5},
	//	tr: helper.Coord{X: 30, Y: -5},
	//	bl: helper.Coord{X: 20, Y: -10},
	//	br: helper.Coord{X: 30, Y: -10},
	//}

	ans, velocities := partOne(box)
	fmt.Printf("Part one: %v\n", ans)
	fmt.Printf("Part two: %v\n", velocities)
}
