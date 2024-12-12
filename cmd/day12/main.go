package main

import (
	"aoc/internal/helpers"
	"bufio"
	"os"
	"strings"
)

func getInput() [][]string {
	f, _ := os.Open("cmd/day12/input.txt")
	scanner := bufio.NewScanner(f)

	plants := make([][]string, 0)
	for scanner.Scan() {
		l := scanner.Text()
		line := strings.Split(l, "")
		plants = append(plants, line)
	}

	return plants
}

type Plant struct {
	X, Y int
	Type string
}

type Fence struct {
	X1, Y1, X2, Y2 int
}

type Garden struct {
	Input [][]string

	Visited map[Plant]bool
	Regions [][]Plant
}

func NewGarden(input [][]string) Garden {
	g := Garden{
		Input:   input,
		Visited: make(map[Plant]bool),
		Regions: make([][]Plant, 0),
	}

	for y, row := range input {
		for x, p := range row {
			if g.Visited[Plant{x, y, p}] {
				continue
			} else {
				g.addRegion(x, y)
			}
		}
	}

	return g
}

func (g *Garden) addRegion(x, y int) {
	plant := Plant{x, y, g.Input[y][x]}

	region := make([]Plant, 0)
	region = append(region, plant)
	g.Visited[plant] = true

	nextNeighbors := g.getNeighbors(x, y)
	for len(nextNeighbors) > 0 {
		temp := make([]Plant, 0)
		for _, n := range nextNeighbors {
			if !g.Visited[n] {
				g.Visited[n] = true
				region = append(region, n)
				temp = append(temp, g.getNeighbors(n.X, n.Y)...)
			}
		}

		nextNeighbors = temp
	}

	g.Regions = append(g.Regions, region)
}

func (g *Garden) getNeighbors(x, y int) []Plant {
	neighbors := make([]Plant, 0)

	plantType := g.Input[y][x]
	if x > 0 && plantType == g.Input[y][x-1] {
		neighbors = append(neighbors, Plant{x - 1, y, g.Input[y][x-1]})
	}
	if x < len(g.Input[y])-1 && plantType == g.Input[y][x+1] {
		neighbors = append(neighbors, Plant{x + 1, y, g.Input[y][x+1]})
	}
	if y > 0 && plantType == g.Input[y-1][x] {
		neighbors = append(neighbors, Plant{x, y - 1, g.Input[y-1][x]})
	}
	if y < len(g.Input)-1 && plantType == g.Input[y+1][x] {
		neighbors = append(neighbors, Plant{x, y + 1, g.Input[y+1][x]})
	}

	return neighbors
}

func (p Plant) fence() []Fence {
	return []Fence{
		{p.X, p.Y, p.X + 1, p.Y},
		{p.X, p.Y, p.X, p.Y + 1},
		{p.X + 1, p.Y, p.X + 1, p.Y + 1},
		{p.X, p.Y + 1, p.X + 1, p.Y + 1},
	}
}

func fence(r []Plant) []Fence {
	fence := make(map[Fence]bool)
	for _, p := range r {
		for _, f := range p.fence() {
			if fence[f] {
				delete(fence, f)
			} else {
				fence[f] = true
			}
		}
	}

	s := make([]Fence, 0)
	for k := range fence {
		s = append(s, k)
	}
	return s
}

func calculateFencePrice(r []Plant) int {
	fenceParts := fence(r)
	return len(fenceParts) * len(r)
}

func (g *Garden) calculateSidePrice(r []Plant) int {
	corners := 0
	for _, p := range r {
		corners += g.corners(p)
	}

	return corners * len(r)
}

func (g *Garden) corners(p Plant) int {
	corners := 0
	if checkNW(p, g) {
		corners++
	}
	if checkNE(p, g) {
		corners++
	}
	if checkSW(p, g) {
		corners++
	}
	if checkSE(p, g) {
		corners++
	}

	return corners
}

func checkNW(p Plant, g *Garden) bool {
	t := p.Type
	if p.Y == 0 {
		return p.X == 0 || t != g.Input[p.Y][p.X-1]
	}

	if p.X == 0 {
		return t != g.Input[p.Y-1][p.X]
	}

	if t != g.Input[p.Y][p.X-1] && t != g.Input[p.Y-1][p.X] {
		return true
	}

	if t == g.Input[p.Y-1][p.X] && t == g.Input[p.Y][p.X-1] && t != g.Input[p.Y-1][p.X-1] {
		return true
	}

	return false
}

func checkNE(p Plant, g *Garden) bool {
	t := p.Type
	if p.Y == 0 {
		return p.X == len(g.Input[p.Y])-1 || t != g.Input[p.Y][p.X+1]
	}

	if p.X == len(g.Input[p.Y])-1 {
		return t != g.Input[p.Y-1][p.X]
	}

	if t != g.Input[p.Y][p.X+1] && t != g.Input[p.Y-1][p.X] {
		return true
	}

	if t == g.Input[p.Y-1][p.X] && t == g.Input[p.Y][p.X+1] && t != g.Input[p.Y-1][p.X+1] {
		return true
	}

	return false
}

func checkSW(p Plant, g *Garden) bool {
	t := p.Type
	if p.Y == len(g.Input)-1 {
		return p.X == 0 || t != g.Input[p.Y][p.X-1]
	}

	if p.X == 0 {
		return t != g.Input[p.Y+1][p.X]
	}

	if t != g.Input[p.Y][p.X-1] && t != g.Input[p.Y+1][p.X] {
		return true
	}

	if t == g.Input[p.Y+1][p.X] && t == g.Input[p.Y][p.X-1] && t != g.Input[p.Y+1][p.X-1] {
		return true
	}

	return false
}

func checkSE(p Plant, g *Garden) bool {
	t := p.Type
	if p.Y == len(g.Input)-1 {
		return p.X == len(g.Input[p.Y])-1 || t != g.Input[p.Y][p.X+1]
	}

	if p.X == len(g.Input[p.Y])-1 {
		return t != g.Input[p.Y+1][p.X]
	}

	if t != g.Input[p.Y][p.X+1] && t != g.Input[p.Y+1][p.X] {
		return true
	}

	if t == g.Input[p.Y+1][p.X] && t == g.Input[p.Y][p.X+1] && t != g.Input[p.Y+1][p.X+1] {
		return true
	}

	return false
}

func solution() (int, int) {
	input := getInput()
	g := NewGarden(input)
	p1 := 0
	p2 := 0
	for _, r := range g.Regions {
		p1 += calculateFencePrice(r)
		p2 += g.calculateSidePrice(r)
	}

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
