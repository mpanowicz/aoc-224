package main

import (
	"aoc/internal/helpers"
	"bufio"
	"math"
	"os"
	"strings"
)

type Puzzle struct {
	Result int
	Values []int
}

func getInput() <-chan Puzzle {
	ch := make(chan Puzzle)
	go func() {
		defer close(ch)

		f, _ := os.Open("cmd/day7/input.txt")
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()
			split := strings.Split(line, ": ")
			result := helpers.ParseInt(split[0])
			values := helpers.ParseInts(strings.Split(split[1], " "))
			puzzle := Puzzle{result, values}
			ch <- puzzle
		}
	}()
	return ch
}

func splitConcatenatedPuzzle(partial int, val int) []int {
	temp := []int{}
	digits := math.Floor(math.Log10(float64(val))) + 1
	div := int(math.Pow10(int(digits)))
	if partial%div == val {
		temp = append(temp, partial/div)
	}

	return temp
}

func splitPuzzle(partials []int, val int, part2 bool) []int {
	temp := []int{}
	for _, v := range partials {
		if v%val == 0 {
			temp = append(temp, v/val)
		}
		if v-val > 0 {
			temp = append(temp, v-val)
		}

		if part2 {
			temp = append(temp, splitConcatenatedPuzzle(v, val)...)
		}
	}

	return temp
}

func correctPuzzle(puzzle Puzzle, part2 bool) bool {
	partials := []int{puzzle.Result}
	for i := len(puzzle.Values) - 1; i > 0; i-- {
		partials = splitPuzzle(partials, puzzle.Values[i], part2)
	}

	for _, v := range partials {
		if v == puzzle.Values[0] {
			return true
		}
	}

	return false
}

func solution() (int, int) {
	sumP1 := 0
	sumP2 := 0
	for puzzle := range getInput() {
		if correctPuzzle(puzzle, false) {
			sumP1 += puzzle.Result
		}
		if correctPuzzle(puzzle, true) {
			sumP2 += puzzle.Result
		}
	}

	return sumP1, sumP2
}

func main() {
	helpers.PrintResult(solution())
}
