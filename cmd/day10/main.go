package main

import (
	"aoc/internal/helpers"
	"bufio"
	"os"
)

type Point struct {
	X int
	Y int

	Value int
}

type TopographicMap struct {
	Points [][]Point
	Starts []Point

	Width  int
	Height int
}

func getInput() TopographicMap {
	f, _ := os.Open("cmd/day10/input.txt")
	scanner := bufio.NewScanner(f)

	points := [][]Point{}
	starts := []Point{}
	y := 0
	for scanner.Scan() {
		l := scanner.Text()
		line := []Point{}
		for x, c := range l {
			val := int(c - '0')
			point := Point{x, y, val}
			if val == 0 {
				starts = append(starts, point)
			}
			line = append(line, point)
		}
		points = append(points, line)
		y++
	}

	return TopographicMap{points, starts, len(points[0]), len(points)}
}

func (tm TopographicMap) findNextPoint(p Point) []Point {
	points := []Point{}
	if p.X+1 < tm.Width {
		next := tm.Points[p.Y][p.X+1]
		if p.Value+1 == next.Value {
			points = append(points, next)
		}
	}
	if p.X-1 >= 0 {
		next := tm.Points[p.Y][p.X-1]
		if p.Value+1 == next.Value {
			points = append(points, next)
		}
	}
	if p.Y+1 < tm.Height {
		next := tm.Points[p.Y+1][p.X]
		if p.Value+1 == next.Value {
			points = append(points, next)
		}
	}
	if p.Y-1 >= 0 {
		next := tm.Points[p.Y-1][p.X]
		if p.Value+1 == next.Value {
			points = append(points, next)
		}
	}
	return points
}

func (tm TopographicMap) findPaths(p Point) (int, int) {
	next := tm.findNextPoint(p)
	for len(next) != 0 && next[0].Value < 9 {
		n := []Point{}
		for _, np := range next {
			n = append(n, tm.findNextPoint(np)...)
		}
		next = n
	}
	distinct := map[Point]struct{}{}
	for _, p := range next {
		distinct[p] = struct{}{}
	}
	return len(distinct), len(next)
}

func solution() (int, int) {
	tm := getInput()
	p1 := 0
	p2 := 0
	for _, s := range tm.Starts {
		x, y := tm.findPaths(s)
		p1 += x
		p2 += y
	}

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
