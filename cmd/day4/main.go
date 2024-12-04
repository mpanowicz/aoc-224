package main

import (
	"aoc/internal/helpers"
	"bufio"
	"os"
)

var XMAS = []byte{'X', 'M', 'A', 'S'}

type Puzzle struct {
	input  [][]byte
	height int
	width  int
}

func NewPuzzle(input [][]byte) Puzzle {
	return Puzzle{
		input,
		len(input),
		len(input[0]),
	}
}

func compare(a []byte) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != XMAS[i] {
			return false
		}
	}

	return true
}

func (p Puzzle) CheckHorizontal(x, y int) int {
	count := 0
	if x+3 < p.width {
		check := []byte{p.input[y][x], p.input[y][x+1], p.input[y][x+2], p.input[y][x+3]}
		if compare(check) {
			count++
		}
	}
	if x-3 >= 0 {
		check := []byte{p.input[y][x], p.input[y][x-1], p.input[y][x-2], p.input[y][x-3]}
		if compare(check) {
			count++
		}
	}
	return count
}

func (p Puzzle) CheckVertical(x, y int) int {
	count := 0
	if y+3 < p.height {
		check := []byte{p.input[y][x], p.input[y+1][x], p.input[y+2][x], p.input[y+3][x]}
		if compare(check) {
			count++
		}
	}
	if y-3 >= 0 {
		check := []byte{p.input[y][x], p.input[y-1][x], p.input[y-2][x], p.input[y-3][x]}
		if compare(check) {
			count++
		}
	}
	return count
}

func (p Puzzle) CheckDiagonal(x, y int) int {
	count := 0
	if x+3 < p.width && y+3 < p.height {
		check := []byte{p.input[y][x], p.input[y+1][x+1], p.input[y+2][x+2], p.input[y+3][x+3]}
		if compare(check) {
			count++
		}
	}
	if x-3 >= 0 && y+3 < p.height {
		check := []byte{p.input[y][x], p.input[y+1][x-1], p.input[y+2][x-2], p.input[y+3][x-3]}
		if compare(check) {
			count++
		}
	}
	if x+3 < p.width && y-3 >= 0 {
		check := []byte{p.input[y][x], p.input[y-1][x+1], p.input[y-2][x+2], p.input[y-3][x+3]}
		if compare(check) {
			count++
		}
	}
	if x-3 >= 0 && y-3 >= 0 {
		check := []byte{p.input[y][x], p.input[y-1][x-1], p.input[y-2][x-2], p.input[y-3][x-3]}
		if compare(check) {
			count++
		}
	}
	return count
}

func (p Puzzle) GetXmasCount() int {
	count := 0
	for y := range p.input {
		for x := range p.input[y] {
			if p.input[y][x] == 'X' {
				count += p.CheckHorizontal(x, y)
				count += p.CheckVertical(x, y)
				count += p.CheckDiagonal(x, y)
			}
		}
	}
	return count
}

func (p Puzzle) CheckCross(x, y int) bool {
	if 0 < x && x < p.width-1 && 0 < y && y < p.height-1 {
		topLeft := p.input[y-1][x-1]
		bottomRight := p.input[y+1][x+1]
		bottomLeft := p.input[y+1][x-1]
		topRight := p.input[y-1][x+1]

		primaryDiagonalCheck := topLeft == 'M' && bottomRight == 'S' || topLeft == 'S' && bottomRight == 'M'
		secondaryDiagonalCheck := bottomLeft == 'M' && topRight == 'S' || bottomLeft == 'S' && topRight == 'M'
		return primaryDiagonalCheck && secondaryDiagonalCheck
	}

	return false
}

func (p Puzzle) GetX_MasCount() int {
	count := 0
	for y := range p.input {
		for x := range p.input[y] {
			if p.input[y][x] == 'A' && p.CheckCross(x, y) {
				count++
			}
		}
	}
	return count
}

func getInput() Puzzle {
	f, _ := os.Open("cmd/day4/input.txt")
	scanner := bufio.NewScanner(f)

	input := [][]byte{}
	for i := 1; scanner.Scan(); i++ {
		input = append(input, []byte(scanner.Text()))
	}
	return NewPuzzle(input)
}

func solution() (int, int) {
	puzzle := getInput()
	p1 := puzzle.GetXmasCount()
	p2 := puzzle.GetX_MasCount()
	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
