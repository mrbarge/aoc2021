package main

import (
	"aoc2021/helper"
	"fmt"
	"os"
)

type Board struct {
	board [5][5]int
	marked [5][5]bool
}

func (b *Board) Print() {
	for x, r := range b.board {
		for y, v := range r {
			fmt.Printf("%v", v)
			if b.marked[x][y] {
				fmt.Printf("* ")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println("")
	}
}

func (b *Board) Winning() bool {
	// check rows
	for row := 0; row < 5; row++ {
		winner := true
		for col := 0; col < 5; col++ {
			if b.marked[row][col] != true {
				winner = false
				break
			}
		}
		if winner {
			return true
		}
	}
	for col := 0; col < 5; col++ {
		winner := true
		for _, row := range b.marked {
			if row[col] != true {
				winner = false
				break
			}
		}
		if winner {
			return true
		}
	}
	return false
}

func (b *Board) Mark(n int) {
	for x, r := range b.board {
		for y, v := range r {
			if v == n {
				b.marked[x][y] = true
			}
		}
	}
}

func (b *Board) Score() int {
	sum := 0
	for x, r := range b.board {
		for y, v := range r {
			if b.marked[x][y] == false {
				sum += v
			}
		}
	}
	return sum
}

func readBoards(lines []string) ([]*Board, error) {
	boards := make([]*Board, 0)

	tmpLines := make([]string, 0)
	for x, v := range lines {
		if v == "" || x == len(lines)-1 {
			if x == len(lines)-1 {
				tmpLines = append(tmpLines, v)
			}
			board, err := readBoard(tmpLines)
			if err != nil {
				return nil, err
			}
			boards = append(boards, board)
			tmpLines = make([]string, 0)
			continue
		}
		tmpLines = append(tmpLines, v)
	}
	return boards, nil
}

func readBoard(lines []string) (*Board, error) {
	b := &Board{}
	for x, l := range lines {
		nums, err := helper.StrCsvToIntArray(l, " ")
		if err != nil {
			return nil, err
		}
		for y, v := range nums {
			b.board[x][y] = v
		}
	}
	return b, nil
}

func partOne(drawnNumbers []int, boards []*Board) (int) {
	for _, drawn := range drawnNumbers {
		for _, board := range boards {
			board.Mark(drawn)
			if board.Winning() {
				return board.Score() * drawn
			}
		}
	}
	return 0
}


func partTwo(drawnNumbers []int, boards []*Board) (int) {
	winners := make([]int, 0)
	for _, drawn := range drawnNumbers {
		for i, board := range boards {
			board.Mark(drawn)
			if board.Winning() {
				foundWinner := false
				for _, w := range winners {
					if i == w {
						foundWinner = true
						break
					}
				}
				if !foundWinner {
					winners = append(winners, i)
				}
			}
		}
		if len(winners) == len(boards) {
			return boards[winners[len(winners)-1]].Score() * drawn
		}
	}
	return 0
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh,false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	drawnNumbers, err := helper.StrCsvToIntArray(lines[0], ",")
	if err != nil {
		fmt.Printf("Unable to read bingo numbers: %v", err)
		return
	}

	boards, err := readBoards(lines[2:])
	if err != nil {
		fmt.Printf("Unable to read bingo numbers: %v", err)
		return
	}

	ans := partOne(drawnNumbers, boards)
	fmt.Printf("Part one: %v\n", ans)

	ans = partTwo(drawnNumbers, boards)
	fmt.Printf("Part two: %v\n", ans)
}
