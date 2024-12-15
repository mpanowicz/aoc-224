package main

import (
	"aoc/internal/helpers"
	"bufio"
	"fmt"
	"os"
	"sort"
)

type FieldType string
type Direction string

var (
	Wall        FieldType = "#"
	Empty       FieldType = "."
	Box         FieldType = "O"
	BigBoxLeft  FieldType = "["
	BigBoxRight FieldType = "]"

	North Direction = "^"
	South Direction = "v"
	West  Direction = "<"
	East  Direction = ">"
)

type Point struct {
	X, Y int
}

type Warehouse struct {
	Fields        [][]FieldType
	Width, Height int

	RobotPosition Point
	Movements     []Direction
	CurrentMove   int
}

func newWarehouse() Warehouse {
	return Warehouse{
		Fields:        [][]FieldType{},
		Width:         0,
		Height:        0,
		RobotPosition: Point{0, 0},
		Movements:     []Direction{},
		CurrentMove:   0,
	}
}

func (w *Warehouse) newBigWarehouse() Warehouse {
	bw := Warehouse{
		Fields:        [][]FieldType{},
		Width:         w.Width * 2,
		Height:        w.Height,
		RobotPosition: Point{w.RobotPosition.X * 2, w.RobotPosition.Y},
		Movements:     w.Movements,
		CurrentMove:   0,
	}

	for y := 0; y < w.Height; y++ {
		l := []FieldType{}
		for x := 0; x < w.Width; x++ {
			if w.Fields[y][x] == Box {
				l = append(l, BigBoxLeft)
				l = append(l, BigBoxRight)
			} else {
				l = append(l, w.Fields[y][x])
				l = append(l, w.Fields[y][x])
			}
		}
		bw.Fields = append(bw.Fields, l)
	}

	return bw
}

func getInput() Warehouse {
	f, _ := os.Open("cmd/day15/input.txt")
	scanner := bufio.NewScanner(f)

	warehouse := newWarehouse()
	readMovements := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			warehouse.Width = len(warehouse.Fields[0])
			readMovements = true
			continue
		}

		if !readMovements {
			warehouse.Height++
			l := []FieldType{}
			for x, c := range line {
				if c == '@' {
					warehouse.RobotPosition = Point{x, warehouse.Height - 1}
					l = append(l, Empty)
				} else {
					l = append(l, FieldType(string(c)))
				}
			}
			warehouse.Fields = append(warehouse.Fields, l)
		} else {
			for _, c := range line {
				warehouse.Movements = append(warehouse.Movements, Direction(string(c)))
			}
		}
	}

	return warehouse
}

func (w *Warehouse) moveRobot() {
	dir := Point{0, 0}
	switch w.Movements[w.CurrentMove] {
	case North:
		dir = Point{0, -1}
	case South:
		dir = Point{0, 1}
	case West:
		dir = Point{-1, 0}
	case East:
		dir = Point{1, 0}
	}
	w.CurrentMove++

	newPos := Point{w.RobotPosition.X + dir.X, w.RobotPosition.Y + dir.Y}
	if w.Fields[newPos.Y][newPos.X] == Empty {
		w.RobotPosition = newPos
	} else if w.Fields[newPos.Y][newPos.X] == Box {
		tempPos := Point{newPos.X + dir.X, newPos.Y + dir.Y}
		for w.Fields[tempPos.Y][tempPos.X] == Box {
			tempPos = Point{tempPos.X + dir.X, tempPos.Y + dir.Y}
		}
		if w.Fields[tempPos.Y][tempPos.X] == Empty {
			w.Fields[tempPos.Y][tempPos.X] = Box
			w.Fields[newPos.Y][newPos.X] = Empty
			w.RobotPosition = newPos
		}
	} else if w.isBigBox(newPos) {
		if dir.X == 0 {
			bigBox := w.getBigBox(newPos)
			bigBoxes := w.getBigBoxes(bigBox, dir)
			if w.canMove(bigBoxes, dir) {
				sort.Slice(bigBoxes, func(i, j int) bool {
					if dir.Y == -1 {
						return bigBoxes[i].Y < bigBoxes[j].Y
					} else {
						return bigBoxes[i].Y > bigBoxes[j].Y
					}
				})
				for _, b := range bigBoxes {
					w.Fields[b.Y+dir.Y][b.X] = w.Fields[b.Y][b.X]
					w.Fields[b.Y][b.X] = Empty
				}
				w.RobotPosition = newPos
			}
		} else {
			tempPos := Point{newPos.X + dir.X, newPos.Y}
			for w.isBigBox(tempPos) {
				tempPos = Point{tempPos.X + dir.X, tempPos.Y}
			}
			if w.Fields[tempPos.Y][tempPos.X] == Empty {
				for {
					if tempPos.X == w.RobotPosition.X {
						break
					}
					w.Fields[tempPos.Y][tempPos.X] = w.Fields[tempPos.Y][tempPos.X-dir.X]
					w.Fields[tempPos.Y][tempPos.X-dir.X] = Empty
					tempPos = Point{tempPos.X - dir.X, tempPos.Y}
				}
				w.RobotPosition = newPos
			}
		}
	}
}

func (w *Warehouse) getBigBox(p Point) []Point {
	bigBox := []Point{p}

	if w.Fields[p.Y][p.X] == BigBoxLeft {
		bigBox = append(bigBox, Point{p.X + 1, p.Y})
	} else {
		bigBox = append(bigBox, Point{p.X - 1, p.Y})
	}

	return bigBox
}

func (w *Warehouse) isBigBox(p Point) bool {
	return w.Fields[p.Y][p.X] == BigBoxLeft || w.Fields[p.Y][p.X] == BigBoxRight
}

func (w *Warehouse) getBigBoxes(box []Point, dir Point) []Point {
	bigBoxes := map[Point]struct{}{}
	for _, b := range box {
		bigBoxes[b] = struct{}{}
	}

	for _, b := range box {
		nextPos := Point{b.X + dir.X, b.Y + dir.Y}
		if w.isBigBox(nextPos) {
			if w.Fields[nextPos.Y][nextPos.X] == w.Fields[b.Y][b.X] {
				if w.Fields[nextPos.Y][nextPos.X] == BigBoxLeft {
					bb := w.getBigBox(nextPos)
					next := w.getBigBoxes(bb, dir)
					for _, n := range next {
						bigBoxes[n] = struct{}{}
					}
				}
			} else {
				bb := w.getBigBox(nextPos)
				next := w.getBigBoxes(bb, dir)
				for _, n := range next {
					bigBoxes[n] = struct{}{}
				}
			}
		}
	}

	boxes := []Point{}
	for b := range bigBoxes {
		boxes = append(boxes, b)
	}

	return boxes
}

func (w *Warehouse) canMove(boxes []Point, dir Point) bool {
	for _, b := range boxes {
		nextPos := Point{b.X + dir.X, b.Y + dir.Y}
		if w.Fields[nextPos.Y][nextPos.X] == Wall {
			return false
		}
	}
	return true
}

func (w *Warehouse) move() {
	for w.CurrentMove < len(w.Movements) {
		w.moveRobot()
	}
}

func (w *Warehouse) print() {
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			if w.RobotPosition.X == x && w.RobotPosition.Y == y {
				fmt.Print("@")
			} else {
				fmt.Print(string(w.Fields[y][x]))
			}
		}
		fmt.Println()
	}
}

func (w *Warehouse) calculateDistance() int {
	sum := 0
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			if w.Fields[y][x] == Box || w.Fields[y][x] == BigBoxLeft {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func solution() (int, int) {
	warehouse := getInput()
	bigWarehouse := warehouse.newBigWarehouse()
	warehouse.move()
	warehouse.print()
	bigWarehouse.move()
	bigWarehouse.print()

	p1 := warehouse.calculateDistance()
	p2 := bigWarehouse.calculateDistance()

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
