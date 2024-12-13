package main

import (
	"aoc/internal/helpers"
	"bufio"
	"os"
	"strings"
)

type Button struct {
	X int
	Y int
}

type Prize struct {
	X int
	Y int
}

type ClawMachine struct {
	A Button
	B Button
	P Prize
}

func parseButton(s string) Button {
	changes := strings.Split(s, ": ")[1]
	coords := strings.Split(changes, ", ")
	x := helpers.ParseInt(strings.Split(coords[0], "+")[1])
	y := helpers.ParseInt(strings.Split(coords[1], "+")[1])
	return Button{x, y}
}

func parsePrize(s string) Prize {
	parts := strings.Split(s, ": ")
	coords := strings.Split(parts[1], ", ")
	x := helpers.ParseInt(strings.Split(coords[0], "=")[1])
	y := helpers.ParseInt(strings.Split(coords[1], "=")[1])
	return Prize{x, y}
}

func getInput() <-chan ClawMachine {
	ch := make(chan ClawMachine)

	go func() {
		defer close(ch)

		f, _ := os.Open("cmd/day13/input.txt")
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			bA := parseButton(scanner.Text())

			scanner.Scan()
			bB := parseButton(scanner.Text())

			scanner.Scan()
			p := parsePrize(scanner.Text())

			ch <- ClawMachine{bA, bB, p}

			scanner.Scan()
		}
	}()

	return ch
}

var (
	tokenA      = 3
	tokenB      = 1
	maxTokens   = 100
	p2prizeDiff = 10000000000000

	part2 = false
)

type Game struct {
	CountA int
	CountB int
}

func (g Game) CalculateTokens() int {
	return g.CountA*tokenA + g.CountB*tokenB
}

func (cm ClawMachine) CheckGame(g Game) bool {
	return cm.A.X*g.CountA+cm.B.X*g.CountB == cm.GetPrize().X && cm.A.Y*g.CountA+cm.B.Y*g.CountB == cm.GetPrize().Y
}

func (cm ClawMachine) CalculateB() int {
	p := cm.GetPrize()
	a := cm.A
	b := cm.B
	count := (p.X*a.Y - p.Y*a.X) / (b.X*a.Y - b.Y*a.X)
	if !part2 && count > maxTokens {
		return maxTokens
	}
	return count
}

func (cm ClawMachine) CalculateA(g Game) Game {
	count := (cm.GetPrize().X - cm.B.X*g.CountB) / cm.A.X
	if !part2 && count > maxTokens {
		g.CountA = maxTokens
	} else {
		g.CountA = count
	}
	return g
}

func (cm ClawMachine) GetPrize() Prize {
	if part2 {
		return Prize{cm.P.X + p2prizeDiff, cm.P.Y + p2prizeDiff}
	} else {
		return cm.P
	}
}

func (cm ClawMachine) CalculateTokens() int {
	maxX := maxTokens * (cm.A.X + cm.B.X)
	maxY := maxTokens * (cm.A.Y + cm.B.Y)
	if !part2 && (maxX < cm.GetPrize().X || maxY < cm.GetPrize().Y) {
		return 0
	}

	countB := cm.CalculateB()
	g := cm.CalculateA(Game{0, countB})

	if cm.CheckGame(g) {
		return g.CalculateTokens()
	} else {
		return 0
	}
}

func solution() (int, int) {
	p1 := 0
	p2 := 0
	for cm := range getInput() {
		part2 = false
		p1 += cm.CalculateTokens()
		part2 = true
		p2 += cm.CalculateTokens()
	}

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
