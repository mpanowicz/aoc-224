package main

import (
	"aoc/internal/helpers"
	"bufio"
	"math"
	"os"
	"strings"
)

func getInput() []int {
	f, _ := os.Open("cmd/day11/input.txt")
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	return helpers.ParseInts(strings.Split(scanner.Text(), " "))
}

var (
	blinkCache = map[int][]int{}
)

func numBlink(v int) []int {
	if v == 0 {
		return []int{1}
	}

	digits := int(math.Floor(math.Log10(float64(v))) + 1)
	if digits%2 == 0 {
		ten := int(math.Pow10(int(digits) / 2))
		right := v % ten
		left := v / ten
		return []int{left, right}
	}

	return []int{v * 2024}
}

func blink(v map[int]int) map[int]int {
	stones := map[int]int{}
	visitCache := map[int]int{}
	for stone, visits := range v {
		if _, ok := visitCache[stone]; !ok {
			blinkCache[stone] = numBlink(stone)
		}
		visitCache[stone] += visits

		for _, nextStone := range blinkCache[stone] {
			stones[nextStone] += visitCache[stone]
		}
	}
	return stones
}

func solution() (int, int) {
	input := getInput()
	temp := map[int]int{}
	for _, v := range input {
		temp[v]++
	}
	for range 25 {
		temp = blink(temp)
	}
	p1 := 0
	for _, visits := range temp {
		p1 += visits
	}
	for range 50 {
		temp = blink(temp)
	}
	p2 := 0
	for _, visits := range temp {
		p2 += visits
	}
	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
