package main

import (
	"aoc/internal/helpers"
	"bufio"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Direction struct {
	X int
	Y int
}

type Robot struct {
	Position  Point
	Direction Direction
}

type Map struct {
	Robots []Robot
}

type Quadrant int

func getInput() Map {
	robots := make([]Robot, 0)

	f, _ := os.Open("cmd/day14/input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")
		pos := strings.Split(parts[0], "=")[1]
		dir := strings.Split(parts[1], "=")[1]

		posVal := strings.Split(pos, ",")
		dirVal := strings.Split(dir, ",")

		r := Robot{
			Point{helpers.ParseInt(posVal[0]), helpers.ParseInt(posVal[1])},
			Direction{helpers.ParseInt(dirVal[0]), helpers.ParseInt(dirVal[1])},
		}
		robots = append(robots, r)
	}

	return Map{robots}
}

var (
	width  = 101
	height = 103

	None        = Quadrant(0)
	TopLeft     = Quadrant(1)
	TopRight    = Quadrant(2)
	BottomRight = Quadrant(3)
	BottomLeft  = Quadrant(4)
)

func mod(a, b int) int {
	return (a%b + b) % b
}

func (r Robot) move(time int) Robot {
	x := mod(r.Position.X+r.Direction.X*time+width, width)
	r.Position.X = x
	y := mod(r.Position.Y+r.Direction.Y*time+height, height)
	r.Position.Y = y

	return r
}

func (r Robot) quadrant() Quadrant {
	wCenter := width / 2
	hCenter := height / 2

	if r.Position.X < wCenter && r.Position.Y < hCenter {
		return TopLeft
	} else if r.Position.X > wCenter && r.Position.Y < hCenter {
		return TopRight
	} else if r.Position.X > wCenter && r.Position.Y > hCenter {
		return BottomRight
	} else if r.Position.X < wCenter && r.Position.Y > hCenter {
		return BottomLeft
	}

	return None
}

func (m Map) Move() {
	for i, r := range m.Robots {
		m.Robots[i] = r.move(1)
	}
}

func (m Map) Quadrant() map[Quadrant]int {
	q := make(map[Quadrant]int)
	for _, r := range m.Robots {
		q[r.quadrant()]++
	}
	return q
}

func countQuadrant(qMap map[Quadrant]int) int {
	v := 1
	for qq, v1 := range qMap {
		if qq != None {
			v *= v1
		}
	}
	return v
}

func checkQuadrant(qMap map[Quadrant]int) bool {
	minCount := 250
	if qMap[TopLeft] > minCount || qMap[TopRight] > minCount || qMap[BottomRight] > minCount || qMap[BottomLeft] > minCount {
		return true
	}
	return false
}

func (m Map) Print() {
	rMap := make(map[Point]bool)
	for _, r := range m.Robots {
		rMap[r.Position] = true
	}
	for y := range height {
		for x := range width {
			if rMap[Point{x, y}] {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
}

func solution() (int, int) {
	m := getInput()
	count := 100
	for range count {
		m.Move()
	}
	qMap := m.Quadrant()
	v := countQuadrant(qMap)
	for {
		count++
		m.Move()
		qMap = m.Quadrant()
		if checkQuadrant(qMap) {
			break
		}
	}

	m.Print()

	return v, count
}

func main() {
	helpers.PrintResult(solution())
}
