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
	posses := []int{1,2,3,4,5,6,7,8,9,10}
	move := dice.roll()
	p.rawpos = (p.rawpos + move) % 10
	p.position = posses[p.rawpos]
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
		fmt.Printf("P1: %v, P2: %v\n",player1.score,player2.score)
	}
	return 0
}

func partTwo(p1 int, p2 int) (int) {

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
		fmt.Printf("P1: %v, P2: %v\n",player1.score,player2.score)
	}
	return 0
}

func main() {
	ans := partOne(8, 1)
	fmt.Printf("Part one: %v\n", ans)
}
