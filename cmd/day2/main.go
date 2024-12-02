package main

import (
	"aoc/internal/helpers"
	"bufio"
	"bytes"
	"os"
)

func getInput() <-chan []int {
	f, _ := os.Open("cmd/day2/input.txt")
	r := bufio.NewReader(f)

	ch := make(chan []int)
	go func() {
		for {
			l, _, _ := r.ReadLine()
			if len(l) == 0 {
				break
			}
			parts := bytes.Split(l, []byte(" "))
			levels := []int{}
			for _, part := range parts {
				levels = append(levels, helpers.ParseInt(string(part)))
			}
			ch <- levels
		}

		close(ch)
	}()

	return ch
}

func checkIncreasing(levels []int) bool {
	increasing := true
	for i := 1; i < len(levels); i++ {
		if levels[i-1] >= levels[i] || levels[i]-levels[i-1] > 3 {
			increasing = false
			break
		}
	}
	return increasing
}

func checkDecreasing(levels []int) bool {
	decreasing := true
	for i := 1; i < len(levels); i++ {
		if levels[i] >= levels[i-1] || levels[i-1]-levels[i] > 3 {
			decreasing = false
			break
		}
	}
	return decreasing
}

func tolerateSingleBadLevel(report []int) bool {
	if checkReport(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		newSlice := make([]int, i)
		copy(newSlice, report[:i])
		part := append(newSlice, report[i+1:]...)
		if checkReport(part) {
			return true
		}
	}

	return false
}

func checkReport(report []int) bool {
	if report[0] < report[1] {
		return checkIncreasing(report)
	} else {
		return checkDecreasing(report)
	}
}

func solution() (int, int) {
	safeCountP1 := 0
	safeCountP2 := 0
	for report := range getInput() {
		if checkReport(report) {
			safeCountP1++
		}
		if tolerateSingleBadLevel(report) {
			safeCountP2++
		}
	}

	return safeCountP1, safeCountP2
}

func main() {
	helpers.PrintResult(solution())
}
