package main

import (
	"aoc/internal/helpers"
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Point struct {
	X, Y int
}

type FieldType string

var (
	Wall  FieldType = "#"
	Empty FieldType = "."

	Cache map[Point]Path = map[Point]Path{}
)

type Maze struct {
	Fields [][]FieldType
	Start  Point
	End    Point
	Width  int
	Height int
}

func (m Maze) Print(p Path) {
	for y, l := range m.Fields {
		for x, f := range l {
			if _, visited := p.Visited[Point{x, y}]; visited {
				fmt.Print("X")
			} else {
				fmt.Print(string(f))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type Path struct {
	Turns       int
	Steps       int
	Direction   Point
	CurrentMove Point
	Visited     map[Point]struct{}
	Path        []Point
}

func createPath(current, dir Point) Path {
	p := Path{
		Turns:       0,
		Steps:       0,
		Direction:   dir,
		CurrentMove: current,
		Visited: map[Point]struct{}{
			current: {},
		},
		Path: []Point{current},
	}
	Cache[current] = p
	return p
}

func (p Point) rotation(n, d Point) (int, Point) {
	dir := Point{n.X - p.X, n.Y - p.Y}
	if dir.X == d.X && dir.Y == d.Y {
		return 0, dir
	} else if math.Abs(float64(dir.X-d.X)) == 2 {
		return 2, dir
	} else if math.Abs(float64(dir.Y-d.Y)) == 2 {
		return 2, dir
	} else {
		return 1, dir
	}
}

func (p Path) extend(next Point) Path {
	v := map[Point]struct{}{
		next: {},
	}
	for k := range p.Visited {
		v[k] = struct{}{}
	}
	path := make([]Point, len(p.Path)+1)
	copy(path, p.Path)
	path[len(path)-1] = next

	r, d := p.CurrentMove.rotation(next, p.Direction)

	return Path{
		Turns:       p.Turns + r,
		Steps:       p.Steps + 1,
		Direction:   d,
		CurrentMove: next,
		Visited:     v,
		Path:        path,
	}
}

func (p Path) calculate() int {
	return p.Turns*1000 + p.Steps
}

func getInput() Maze {
	m := Maze{
		Fields: [][]FieldType{},
		Start:  Point{0, 0},
		End:    Point{0, 0},
		Width:  0,
		Height: 0,
	}

	f, _ := os.Open("cmd/day16/input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		l := []FieldType{}
		for x, c := range scanner.Text() {
			if c == 'S' {
				m.Start = Point{x, m.Height}
				l = append(l, Empty)
			} else if c == 'E' {
				m.End = Point{x, m.Height}
				l = append(l, Empty)
			} else {
				l = append(l, FieldType(string(c)))
			}
		}
		m.Fields = append(m.Fields, l)
		m.Height++
	}
	m.Width = len(m.Fields[0])

	return m
}

func (m Maze) isWall(p Point) bool {
	return m.Fields[p.Y][p.X] == Wall
}

func (m Maze) getNext(p Point) []Point {
	ps := []Point{}
	if p.X > 0 {
		n := Point{p.X - 1, p.Y}
		if !m.isWall(n) {
			ps = append(ps, n)
		}
	}
	if p.X < m.Width-1 {
		n := Point{p.X + 1, p.Y}
		if !m.isWall(n) {
			ps = append(ps, n)
		}
	}
	if p.Y > 0 {
		n := Point{p.X, p.Y - 1}
		if !m.isWall(n) {
			ps = append(ps, n)
		}
	}
	if p.Y < m.Height-1 {
		n := Point{p.X, p.Y + 1}
		if !m.isWall(n) {
			ps = append(ps, n)
		}
	}
	return ps
}

func (m Maze) findPath(p Path) []Path {
	next := m.getNext(p.CurrentMove)
	paths := []Path{}
	for _, n := range next {
		if _, ok := p.Visited[n]; !ok {
			extended := p.extend(n)
			if cached, ok := Cache[n]; ok {
				if cached.Turns+1 < extended.Turns {
					continue
				}
			} else {
				Cache[n] = extended
			}
			paths = append(paths, extended)
		}
	}

	return paths
}

func (m Maze) findEnd() []Path {
	path := createPath(m.Start, Point{1, 0})
	paths := m.findPath(path)
	ends := []Path{}
	for {
		if len(paths) == 0 {
			break
		}

		p := paths[0]
		if p.CurrentMove == m.End {
			ends = append(ends, p)
			paths = paths[1:]
		} else {
			paths = append(paths[1:], m.findPath(p)...)
			sort.Slice(paths, func(i, j int) bool {
				return paths[i].calculate() < paths[j].calculate()
			})
		}
	}

	sort.Slice(ends, func(i, j int) bool {
		return ends[i].calculate() < ends[j].calculate()
	})

	shortest := []Path{}
	c := ends[0].calculate()
	for _, e := range ends {
		if e.calculate() == c {
			shortest = append(shortest, e)
		}
	}

	return shortest
}

func solution() (int, int) {
	maze := getInput()
	paths := maze.findEnd()
	maze.Print(paths[0])
	p1 := paths[0].calculate()

	distinct := map[Point]struct{}{}
	for _, p := range paths {
		for _, v := range p.Path {
			distinct[v] = struct{}{}
		}
	}
	p2 := len(distinct)

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
