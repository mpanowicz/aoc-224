package main

import (
	"aoc/internal/helpers"
	"bufio"
	"os"
)

type Input struct {
	Locks [][]int
	Keys  [][]int
}

func getInput() Input {
	f, _ := os.Open("cmd/day25/input.txt")
	scanner := bufio.NewScanner(f)

	i := Input{[][]int{}, [][]int{}}
	scanType := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		l := scanner.Text()
		if scanType == 0 {
			if l == "....." {
				scanType = 1 //key
			} else {
				scanType = -1 //lock
			}
		}
		parts := [][]rune{}
		for range 5 {
			scanner.Scan()
			parts = append(parts, []rune(scanner.Text()))
		}
		scanner.Scan()
		if scanType == 1 {
			vals := []int{}
			for x := 0; x < 5; x++ {
				sum := 0
				for y := 4; y >= 0; y-- {
					if parts[y][x] == '#' {
						sum++
					} else {
						break
					}
				}
				vals = append(vals, sum)
			}
			i.Keys = append(i.Keys, vals)
		} else {
			vals := []int{}
			for x := 0; x < 5; x++ {
				sum := 0
				for y := 0; y < 5; y++ {
					if parts[y][x] == '#' {
						sum++
					} else {
						break
					}
				}
				vals = append(vals, sum)
			}
			i.Locks = append(i.Locks, vals)
		}
		scanType = 0
	}
	return i
}

func solution() (int, int) {
	in := getInput()
	sum := 0
	for _, k := range in.Keys {
		for _, l := range in.Locks {
			ok := true
			for i := range k {
				if k[i]+l[i] > 5 {
					ok = false
					break
				}
			}
			if ok {
				sum++
			}
		}
	}
	return sum, 0
}

func main() {
	helpers.PrintResult(solution())
}
