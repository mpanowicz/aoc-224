package main

import (
	"aoc/internal/helpers"
	"bufio"
	"fmt"
	"math"
	"os"
)

type Secret struct {
	Start   int
	End     int
	Changes []int
	Diffs   map[string]int
}

func getInput() []int {
	nums := []int{}
	f, _ := os.Open("cmd/day22/input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		nums = append(nums, helpers.ParseInt(scanner.Text()))
	}
	return nums
}

func mix(x, y int) int {
	return (x ^ y) % 16777216
}

func calculate(n int) int {
	next := mix(n, n*64)
	next = mix(next, next/32)
	next = mix(next, next*2048)
	return next
}

func calculateN(num int, n int, cache *map[string]int) Secret {
	s := Secret{Start: num, End: 0, Changes: []int{}, Diffs: map[string]int{}}
	temp := num
	for i := 0; i < n; i++ {
		num = calculate(temp)
		dif := (num%10 - temp%10)
		s.Changes = append(s.Changes, dif)
		temp = num

		if i > 2 {
			seq := fmt.Sprintf("%d%d%d%d", s.Changes[i-3], s.Changes[i-2], s.Changes[i-1], s.Changes[i])
			if _, ok := s.Diffs[seq]; !ok {
				s.Diffs[seq] = num % 10
				(*cache)[seq] += num % 10
			}
		}
	}
	s.End = num
	return s
}

func solution() (int, int) {
	nums := getInput()
	p1 := 0
	cache := map[string]int{}
	for _, num := range nums {
		s := calculateN(num, 2000, &cache)
		p1 += s.End
	}
	p2 := math.MinInt
	for _, v := range cache {
		if v > p2 {
			p2 = v
		}
	}
	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
