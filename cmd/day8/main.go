package main

import (
	"aoc/internal/helpers"
	"bufio"
	"math"
	"os"
)

type Antenna struct {
	X int
	Y int
}

type Puzzle struct {
	Antennas map[rune][]Antenna
	Width    int
	Height   int
}

func getInput() Puzzle {
	f, _ := os.Open("cmd/day8/input.txt")
	scanner := bufio.NewScanner(f)

	y := 0
	x := 0
	antennas := make(map[rune][]Antenna)
	for scanner.Scan() {
		line := scanner.Text()
		if y == 0 {
			x = len(line)
		}

		for x, c := range line {
			if c != '.' {
				antennas[c] = append(antennas[c], Antenna{x, y})
			}
		}
		y++
	}

	return Puzzle{antennas, x, y}
}

func (a Antenna) OnMap(p Puzzle) bool {
	return a.X >= 0 && a.X < p.Width && a.Y >= 0 && a.Y < p.Height
}

func (p Puzzle) GenerateAntinodes(part2 bool) map[Antenna]struct{} {
	antinodes := make(map[Antenna]struct{})
	for _, antennas := range p.Antennas {
		for i := 0; i < len(antennas); i++ {
			a := antennas[i]
			for j := i + 1; j < len(antennas); j++ {
				b := antennas[j]

				yDistance := int(math.Abs(float64(b.Y - a.Y)))
				xDistance := int(math.Abs(float64(b.X - a.X)))

				if b.Y < a.Y {
					if b.X > a.X {
						antinode := Antenna{b.X + xDistance, b.Y - yDistance}
						for {
							if antinode.OnMap(p) {
								antinodes[antinode] = struct{}{}
								antinode = Antenna{antinode.X + xDistance, antinode.Y - yDistance}
							} else {
								break
							}
							if !part2 {
								break
							}
						}

						antinode = Antenna{a.X - xDistance, a.Y + yDistance}
						for {
							if antinode.OnMap(p) {
								antinodes[antinode] = struct{}{}
								antinode = Antenna{antinode.X - xDistance, antinode.Y + yDistance}
							} else {
								break
							}
							if !part2 {
								break
							}
						}
					} else {
						antinode := Antenna{b.X - xDistance, b.Y - yDistance}
						for {
							if antinode.OnMap(p) {
								antinodes[antinode] = struct{}{}
								antinode = Antenna{antinode.X - xDistance, antinode.Y - yDistance}
							} else {
								break
							}
							if !part2 {
								break
							}
						}

						antinode = Antenna{a.X + xDistance, a.Y + yDistance}
						for {
							if antinode.OnMap(p) {
								antinodes[antinode] = struct{}{}
								antinode = Antenna{antinode.X + xDistance, antinode.Y + yDistance}
							} else {
								break
							}
							if !part2 {
								break
							}
						}
					}
				} else {
					if b.X > a.X {
						antinode := Antenna{b.X + xDistance, b.Y + yDistance}
						for {
							if antinode.OnMap(p) {
								antinodes[antinode] = struct{}{}
								antinode = Antenna{antinode.X + xDistance, antinode.Y + yDistance}
							} else {
								break
							}
							if !part2 {
								break
							}
						}

						antinode = Antenna{a.X - xDistance, a.Y - yDistance}
						for {
							if antinode.OnMap(p) {
								antinodes[antinode] = struct{}{}
								antinode = Antenna{antinode.X - xDistance, antinode.Y - yDistance}
							} else {
								break
							}
							if !part2 {
								break
							}
						}
					} else {
						antinode := Antenna{b.X - xDistance, b.Y + yDistance}
						for {
							if antinode.OnMap(p) {
								antinodes[antinode] = struct{}{}
								antinode = Antenna{antinode.X - xDistance, antinode.Y + yDistance}
							} else {
								break
							}
							if !part2 {
								break
							}
						}

						antinode = Antenna{a.X + xDistance, a.Y - yDistance}
						for {
							if antinode.OnMap(p) {
								antinodes[antinode] = struct{}{}
								antinode = Antenna{antinode.X + xDistance, antinode.Y - yDistance}
							} else {
								break
							}
							if !part2 {
								break
							}
						}
					}
				}
				if part2 {
					antinodes[a] = struct{}{}
					antinodes[b] = struct{}{}
				}
			}
		}
	}

	return antinodes
}

func solution() (int, int) {
	puzzle := getInput()
	p1 := len(puzzle.GenerateAntinodes(false))
	p2 := len(puzzle.GenerateAntinodes(true))

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
