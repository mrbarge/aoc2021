package main

import (
	"fmt"
)

type Player struct {
	position int
	rawpos int
	score    int
}

type Dice struct {
	next int
	rolls int
}

var dice = &Dice{next: 1}
var games = make(map[string][2]int64)
var dicenums = []int{1,2,3,4,5,6,7,8,9,10}

func (d *Dice) roll() int {
	r := 0
	for i := 0; i < 3; i++ {
		r += d.next
		d.next++
		if d.next > 100 {
			d.next = 1
		}
	}
	d.rolls += 3
	return r
}

func (p *Player) turn() {
	move := dice.roll()
	p.rawpos = (p.rawpos + move) % 10
	p.position = dicenums[p.rawpos]
	p.score += p.position
}

func partOne(p1 int, p2 int) (int) {

	player1 := &Player{position: p1, rawpos:p1-1}
	player2 := &Player{position: p2, rawpos:p2-1}

	dice.next = 1
	done := false
	for !done {
		player1.turn()
		if player1.score >= 1000 {
			done = true
			return player2.score * dice.rolls
		}
		player2.turn()
		if player2.score >= 1000 {
			done = true
			return player1.score * dice.rolls
		}
	}
	return 0
}

func partTwo(p1 int, p2 int) int64 {
	v := turn(3, p1, p2, 0, 0)
	if v[0] > v[1] {
		return v[0]
	} else {
		return v[1]
	}
}

func turn(rolls int, p1 int, p2 int, p1score int, p2score int) [2]int64 {
	t := fmt.Sprintf("%v,%v,%v,%v,%v",rolls,p1,p2,p1score,p2score)
	if v, ok := games[t]; ok {
		return v
	}

	p1Wins := int64(0)
	p2Wins := int64(0)

	if rolls == 0 {
		p1score += p1
		if p1score >= 21 {
			return [2]int64{1,0}
		}
		result := turn(3, p2, p1, p2score, p1score)
		return [2]int64{result[1],result[0]}
	}

	for _, roll := range []int{1,2,3} {
		p1NextPos := (p1 + roll) % 10
		if p1NextPos == 0 {
			p1NextPos = 10
		}
		results := turn(rolls-1,p1NextPos,p2,p1score,p2score)
		p1Wins += results[0]
		p2Wins += results[1]
	}

	r := [2]int64{p1Wins, p2Wins}
	games[t] = r
	return r
}

func main() {
	ans := partOne(8, 1)
	fmt.Printf("Part one: %v\n", ans)
	ans2 := partTwo(8, 1)
	fmt.Printf("Part two: %v\n", ans2)
}
