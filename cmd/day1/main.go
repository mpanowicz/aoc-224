package main

import (
	"aoc/internal/helpers"
	"bufio"
	"bytes"
	"math"
	"os"
	"slices"
)

func getInput() ([]int, []int) {
	f, _ := os.Open("cmd/day1/input.txt")
	r := bufio.NewReader(f)

	left := []int{}
	right := []int{}
	for {
		line, _, _ := r.ReadLine()
		if len(line) == 0 {
			break
		}
		parts := bytes.Split(line, ([]byte)("   "))
		l := helpers.ParseInt(string(parts[0]))
		r := helpers.ParseInt(string(parts[1]))
		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}

func part1(left []int, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	diff := 0
	for i := 0; i < len(left); i++ {
		diff += int(math.Abs(float64(left[i]) - float64(right[i])))
	}
	return diff
}

func part2(left []int, right []int) int {
	leftMap := map[int]int{}
	rightMap := map[int]int{}

	for _, l := range left {
		if v, ok := leftMap[l]; ok {
			leftMap[l] = v + 1
		} else {
			leftMap[l] = 1
		}
	}
	for _, r := range right {
		if v, ok := rightMap[r]; ok {
			rightMap[r] = v + 1
		} else {
			rightMap[r] = 1
		}
	}

	diff := 0
	for _, v := range left {
		if _, ok := leftMap[v]; ok {
			if r, ok := rightMap[v]; ok {
				diff += v * r
			}
		}
	}

	return diff
}

func solution() (int, int) {
	left, right := getInput()

	p1 := part1(left, right)
	p2 := part2(left, right)

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
