package main

import (
	"aoc/internal/helpers"
	"bufio"
	"fmt"
	"os"
)

type Field string

type Point struct {
	X, Y int
}

var (
	Empty Field = "."
	Wall  Field = "#"

	Directions = []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

type Move struct {
	Point Point
	Steps int
}

type RaceTrack struct {
	Fields [][]Field
	Start  Point
	End    Point

	Width  int
	Height int

	Track    []Move
	TrackMap map[Point]int
}

type Cheat struct {
	From Point
	To   Point
}

func getInput() RaceTrack {
	f, _ := os.Open("cmd/day20/input.txt")
	scanner := bufio.NewScanner(f)

	rt := RaceTrack{
		Fields: make([][]Field, 0),
	}
	for scanner.Scan() {
		line := scanner.Text()
		if rt.Width == 0 {
			rt.Width = len(line)
		}
		row := make([]Field, 0)
		for x, c := range line {
			switch c {
			case 'S':
				row = append(row, Empty)
				rt.Start = Point{x, rt.Height}
			case 'E':
				row = append(row, Empty)
				rt.End = Point{x, rt.Height}
			case '.':
				row = append(row, Empty)
			case '#':
				row = append(row, Wall)
			}
		}
		rt.Fields = append(rt.Fields, row)
		rt.Height++
	}
	return rt
}

func (rt *RaceTrack) validPoint(p Point) bool {
	return p.X >= 0 && p.X < rt.Width && p.Y >= 0 && p.Y < rt.Height
}

func (rt *RaceTrack) findTrack() {
	rt.Track = make([]Move, 0)
	rt.TrackMap = make(map[Point]int)

	visited := make(map[Point]bool)
	visited[rt.Start] = true
	current := rt.Start
	for current != rt.End {
		for _, dir := range Directions {
			next := Point{current.X + dir.X, current.Y + dir.Y}
			if rt.Fields[next.Y][next.X] == Wall || visited[next] {
				continue
			} else {
				steps := len(rt.Track) + 1
				rt.Track = append(rt.Track, Move{next, steps})
				rt.TrackMap[next] = steps
				visited[next] = true
				current = next
				break
			}
		}
	}
}

func (rt *RaceTrack) findSingleWallCheats() map[int]int {
	res := map[int]int{}
	visited := make(map[Cheat]bool)

	for _, move := range rt.Track {
		for _, dir := range Directions {
			next := Point{move.Point.X + dir.X*2, move.Point.Y + dir.Y*2}
			expectedWall := Point{move.Point.X + dir.X, move.Point.Y + dir.Y}
			c := Cheat{move.Point, next}
			if !rt.validPoint(next) || rt.Fields[next.Y][next.X] == Wall || visited[c] || rt.Fields[expectedWall.Y][expectedWall.X] != Wall {
				continue
			} else {
				cheatVal := rt.TrackMap[next]
				diff := cheatVal - move.Steps - 2
				if diff > 0 {
					res[diff]++
					visited[c] = true
					fmt.Println(c, rt.TrackMap[c.From], rt.TrackMap[c.To], diff)
					visited[Cheat{next, move.Point}] = true
				}
			}
		}
	}

	return res
}

func manhattan(p1, p2 Point) int {
	return helpers.Abs(p1.X-p2.X) + helpers.Abs(p1.Y-p2.Y)
}

func (rt *RaceTrack) findLongCheats() map[int]int {
	res := map[int]int{}
	starts := []Move{{rt.Start, 0}}
	starts = append(starts, rt.Track[:len(rt.Track)-1]...)

	for _, start := range starts {
		for i := len(rt.Track) - 1; i >= 0; i-- {
			m := manhattan(start.Point, rt.Track[i].Point)
			if m <= 20 {
				cheatVal := rt.TrackMap[rt.Track[i].Point]
				diff := cheatVal - start.Steps - m
				if diff > 0 {
					res[diff]++
				}
			}
		}
	}

	return res
}

func solution() (int, int) {
	rt := getInput()
	rt.findTrack()
	cheats := rt.findSingleWallCheats()
	fmt.Println(cheats)
	p1 := 0
	for k, v := range cheats {
		if k >= 100 {
			p1 += v
		}
	}
	cheats = rt.findLongCheats()
	p2 := 0
	for k, v := range cheats {
		if k >= 100 {
			p2 += v
		}
	}

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
