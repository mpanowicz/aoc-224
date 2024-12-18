package main

import (
	"aoc/internal/helpers"
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Memory struct {
	Width     int
	Height    int
	Corrupted map[Point]struct{}

	NextCorrupted []Point
	CorruptedIdx  int
}

var (
	Cache = make(map[Point]int)
)

func (m Memory) Print(p Path) {
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			coords := Point{x, y}
			if _, corrupted := m.Corrupted[coords]; corrupted {
				fmt.Print("#")
			} else if _, visited := p.Visited[coords]; visited {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type Path struct {
	Visited map[Point]struct{}
	Points  []Point
}

func getInput() Memory {
	s := 71
	m := Memory{Width: s, Height: s}

	f, _ := os.Open("cmd/day18/input.txt")
	scanner := bufio.NewScanner(f)

	m.NextCorrupted = []Point{}
	m.Corrupted = make(map[Point]struct{})
	for scanner.Scan() {
		l := scanner.Text()
		coords := strings.Split(l, ",")
		x := helpers.ParseInt(coords[0])
		y := helpers.ParseInt(coords[1])
		if len(m.Corrupted) < 1024 {
			m.Corrupted[Point{x, y}] = struct{}{}
		} else {
			m.NextCorrupted = append(m.NextCorrupted, Point{x, y})
		}
	}
	return m
}

func (m Memory) findNext(p Point) []Point {
	next := []Point{}
	if p.X > 0 {
		p := Point{p.X - 1, p.Y}
		if _, corrupted := m.Corrupted[p]; !corrupted {
			next = append(next, p)
		}
	}
	if p.X < m.Width-1 {
		p := Point{p.X + 1, p.Y}
		if _, corrupted := m.Corrupted[p]; !corrupted {
			next = append(next, p)
		}
	}
	if p.Y > 0 {
		p := Point{p.X, p.Y - 1}
		if _, corrupted := m.Corrupted[p]; !corrupted {
			next = append(next, p)
		}
	}
	if p.Y < m.Height-1 {
		p := Point{p.X, p.Y + 1}
		if _, corrupted := m.Corrupted[p]; !corrupted {
			next = append(next, p)
		}
	}
	return next
}

func (m Memory) findPath() Path {
	s := Point{0, 0}
	e := Point{m.Width - 1, m.Height - 1}

	p := Path{Visited: map[Point]struct{}{s: {}}, Points: []Point{s}}
	paths := []Path{p}
	for len(paths) > 0 {
		temp := []Path{}
		for _, path := range paths {
			last := path.Points[len(path.Points)-1]
			if cv, cached := Cache[last]; cached {
				if cv < len(path.Points)-1 {
					continue
				}
			}

			if last == e {
				return path
			}
			next := m.findNext(last)

			for _, n := range next {
				if _, visited := path.Visited[n]; !visited {

					if _, cached := Cache[n]; cached {
						continue
					}
					Cache[n] = len(path.Points) + 1

					visited := map[Point]struct{}{n: {}}
					for k := range path.Visited {
						visited[k] = struct{}{}
					}
					points := make([]Point, len(path.Points)+1)
					copy(points, path.Points)
					points[len(points)-1] = n
					temp = append(temp, Path{Visited: visited, Points: points})
				}
			}
		}
		paths = temp
		sort.Slice(paths, func(i, j int) bool {
			lastI := paths[i].Points[len(paths[i].Points)-1]
			lastJ := paths[j].Points[len(paths[j].Points)-1]
			distI := math.Abs(float64(lastI.X-e.X)) + math.Abs(float64(lastI.Y-e.Y))
			distJ := math.Abs(float64(lastJ.X-e.X)) + math.Abs(float64(lastJ.Y-e.Y))
			return distI < distJ
		})
	}

	return Path{}
}

func (m *Memory) addAllCorrupted() {
	for _, c := range m.NextCorrupted {
		m.Corrupted[c] = struct{}{}
		m.CorruptedIdx++
	}
}

func (m *Memory) corruptRemove() {
	if m.CorruptedIdx > 0 {
		m.CorruptedIdx--
		delete(m.Corrupted, m.NextCorrupted[m.CorruptedIdx])
	}
}

func solution() (int, string) {
	m := getInput()
	path1 := m.findPath()
	m.Print(path1)
	p1 := len(path1.Points) - 1

	m.addAllCorrupted()
	for {
		Cache = make(map[Point]int)
		res := m.findPath()
		if len(res.Points) > 0 {
			m.Print(res)
			break
		} else {
			fmt.Println(m.CorruptedIdx)
		}
		m.corruptRemove()
	}
	res := m.NextCorrupted[m.CorruptedIdx]
	p2 := fmt.Sprintf("%d,%d", res.X, res.Y)

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
