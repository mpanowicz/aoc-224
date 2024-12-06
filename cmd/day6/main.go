package main

import (
	"aoc/internal/helpers"
	"bufio"
	"os"
)

type Point struct {
	x, y int
}

type Status string
type Direction string

const (
	Visited Status = "visited"
	Blocked Status = "blocked"
	Empty   Status = "empty"

	Up    Direction = "up"
	Down  Direction = "down"
	Left  Direction = "left"
	Right Direction = "right"
)

type Grid struct {
	Points    map[Point]Status
	Current   Point
	Direction Direction
	Width     int
	Height    int

	visited    int
	TurnsCount int
}

func getInput() Grid {
	f, _ := os.Open("cmd/day6/input.txt")
	scanner := bufio.NewScanner(f)

	width := 0
	height := 0
	points := make(map[Point]Status)
	current := Point{0, 0}
	for scanner.Scan() {
		line := scanner.Text()

		if width == 0 {
			width = len(line)
		}

		for i, c := range line {
			if c == '#' {
				points[Point{i, height}] = Blocked
			} else if c == '^' {
				points[Point{i, height}] = Visited
				current = Point{i, height}
			} else {
				points[Point{i, height}] = Empty
			}
		}

		height++
	}

	return Grid{points, current, "up", width, height, 1, 0}
}

func (d Direction) turnRight() Direction {
	switch d {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	}

	return Up
}

func (g *Grid) move() bool {
	next := Point{g.Current.x, g.Current.y}
	switch g.Direction {
	case Up:
		next.y--
	case Down:
		next.y++
	case Left:
		next.x--
	case Right:
		next.x++
	}

	if g.Points[next] == Blocked {
		g.Direction = g.Direction.turnRight()
		g.TurnsCount++
	} else {
		g.Current = next
		if val, ok := g.Points[next]; ok {
			if val == Empty {
				g.Points[next] = Visited
				g.visited++
			}
		} else {
			return false
		}
	}
	return true
}

func (g Grid) cycle() bool {
	for g.move() {
		if g.TurnsCount > g.Height*g.Width {
			return true
		}
	}
	return false
}

func (g Grid) copy() Grid {
	pointsCopy := make(map[Point]Status)
	for k, v := range g.Points {
		pointsCopy[k] = v
	}
	return Grid{pointsCopy, g.Current, g.Direction, g.Width, g.Height, g.visited, g.TurnsCount}
}

func solvePart1(g Grid) int {
	copy := g.copy()
	for copy.move() {
	}
	return copy.visited
}

func solvePart2(g Grid) int {
	count := 0

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if g.Points[Point{x, y}] == Empty {
				copy := g.copy()
				copy.Points[Point{x, y}] = Blocked
				if copy.cycle() {
					count++
				}
			}
		}
	}

	return count
}

func solution() (int, int) {
	grid := getInput()
	p1 := solvePart1(grid)
	p2 := solvePart2(grid)

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
