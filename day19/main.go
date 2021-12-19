package main

import (
	"aoc2021/helper"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Scanner struct {
	id       int
	beacons  []Beacon
	position helper.Coord3D
}

type Beacon [3]int

var posMaps = [][3]int{
	{0, 1, 2},
	{0, 2, 1},
	{1, 0, 2},
	{1, 2, 0},
	{2, 1, 0},
	{2, 0, 1},
}

var negMaps = [][3]int{
	{1, 1, 1},
	{1, 1, -1},
	{1, -1, 1},
	{1, -1, -1},
	{-1, 1, 1},
	{-1, 1, -1},
	{-1, -1, 1},
	{-1, -1, -1},
}

func allSet(scanners []*Scanner) bool {
	l := len(scanners)
	c := 0
	for _, s := range scanners {
		if s != nil {
			c++
		}
	}
	return c == l
}

func partOne(scanners []*Scanner) (int, int, error) {

	aligned := make([]*Scanner, len(scanners))
	aligned[0] = scanners[0]
	failed := make(map[string]bool, 0)
	beaconmap := make(map[string]bool, 0)
	allScandiffs := make([][3]int, 0)
	for _, sv := range scanners[0].beacons {
		beaconmap[fmt.Sprintf("%s-%s-%s",sv[0],sv[1],sv[2])] = true
	}

	for !allSet(aligned) {
		for x, scanner := range scanners {
			if aligned[x] != nil {
				continue
			}
			for y, alignedScanner := range aligned {
				if alignedScanner == nil {
					continue
				}
				if x == y {
					continue
				}
				if failed[fmt.Sprintf("%s-%s",x,y)] {
					continue
				}
				if failed[fmt.Sprintf("%s-%s",x,y)] {
					continue
				}
				fmt.Printf("Comparing %v to %v\n",scanner.id,alignedScanner.id)
				rb, scandiff := alignedScanner.relativeBeacons(scanner)
				if len(rb) > 0 {
					fmt.Printf("Aligned %v to %v",x,y)
					aligned[x] = &Scanner{
						id: x,
					}
					for _, v := range rb {
						aligned[x].beacons = append(aligned[x].beacons, Beacon{v[0],v[1],v[2]})
						beaconmap[fmt.Sprintf("%s-%s-%s",v[0],v[1],v[2])] = true
					}
					allScandiffs = append(allScandiffs, scandiff)
				} else {
					failed[fmt.Sprintf("%s-%s",x,y)] = true
					failed[fmt.Sprintf("%s-%s",y,x)] = true
				}
			}
		}
	}
	max := math.MinInt64
	for _, diff := range allScandiffs {
		for _, diff2 := range allScandiffs {
			dist := manhattan(diff, diff2)
			if dist > max {
				max = dist
			}
		}
	}
	return len(beaconmap), max, nil
}

func manhattan(d1 [3]int, d2 [3]int) int {
	s := math.Abs(float64(d2[0]-d1[0]))  +  math.Abs(float64(d2[1]-d1[1])) + math.Abs(float64(d2[2]-d1[2]))
	return int(s)
}

func relative(b1 [3]int, b2 [3]int) [3]int {
	return [3]int{
		b2[0]-b1[0],
		b2[1]-b1[1],
		b2[2]-b1[2],
	}
}

func (s *Scanner) relativeBeacons(t *Scanner) ([][3]int, [3]int) {

	for _, pm := range posMaps {
		for _, nm := range negMaps {

			scanperms := make([][3]int, 0)
			for _, beacon := range t.beacons {
				scanperms = append(scanperms, [3]int{
					beacon[pm[0]]*nm[0],
					beacon[pm[1]]*nm[1],
					beacon[pm[2]]*nm[2],
				})
			}

			for _, beacon := range s.beacons {
				for _, coord := range scanperms {
					matches := 0
					remap_beacons := make([][3]int, 0)
					// determine relativity to s
					sdiff := relative(beacon,coord)
					// now check all potential matches using this permutation
					for _, matchperm := range scanperms {
						// compute relative
						mrel := relative(sdiff, matchperm)
						for _, bmatch := range s.beacons {
							if bmatch[0] == mrel[0] && bmatch[1] == mrel[1] && bmatch[2] == mrel[2] {
								matches++
							}
						}
						remap_beacons = append(remap_beacons, mrel)
					}
					if matches >= 12 {
						return remap_beacons, sdiff
					}
				}
			}

		}
	}
	return [][3]int{}, [3]int{}
}

func coordPermutations(p [3]int) [][3]int {
	r := make([][3]int, 0)
	for _, pm := range posMaps {
		for _, nm := range negMaps {
			r = append(r, [3]int{
				p[pm[0]]*nm[0],
				p[pm[1]]*nm[1],
				p[pm[2]]*nm[2],
			})
		}
	}
	return r
}

func readData(lines []string) []*Scanner {
	scanners := make([]*Scanner, 0)
	currentScanner := &Scanner{id: 0}
	currentBeacons := make([]Beacon, 0)
	scannerCount := 0

	for _, line := range lines {
		if strings.Contains(line, "---") {
			if len(currentBeacons) > 0 {
				// wrap up existing scanner
				currentScanner.beacons = currentBeacons
				scanners = append(scanners, currentScanner)
				scannerCount++
				currentScanner = &Scanner{id:scannerCount}
				currentBeacons = make([]Beacon, 0)
			}
		} else {
			elems := strings.Split(line, ",")
			x, _ := strconv.Atoi(elems[0])
			y, _ := strconv.Atoi(elems[1])
			z, _ := strconv.Atoi(elems[2])
			currentBeacons = append(currentBeacons, Beacon{x,y,z})
		}
	}
	if len(currentBeacons) > 0 {
		// wrap up existing scanner
		currentScanner.beacons = currentBeacons
		scanners = append(scanners, currentScanner)
	}
	return scanners
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	scanners := readData(lines)
	ans, p2, err := partOne(scanners)
	fmt.Printf("Part one: %v\n", ans)
	fmt.Printf("Part two: %v", p2)
}
