package main

import (
	"aoc/internal/helpers"
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	log = false
)

func getInput() []Number {
	f, _ := os.Open("cmd/day9/input.txt")
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	input := scanner.Text()

	numbers := []Number{}
	for i := 0; i < len(input); i++ {
		if i%2 == 0 {
			numbers = append(numbers, Number{i / 2, int(input[i] - '0')})
		} else {
			numbers = append(numbers, Number{-1, int(input[i] - '0')})
		}
	}

	return numbers
}

type Number struct {
	Value int
	Len   int
}

func print(numbers []Number) {
	if !log {
		return
	}

	for _, n := range numbers {
		for range n.Len {
			if n.Value == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(n.Value)
			}
		}
	}
	fmt.Println()
}

func sum(numbers []Number) int {
	sum := 0
	i := 0
	for _, m := range numbers {
		for range m.Len {
			if m.Value > 0 {
				sum += i * m.Value
			}
			i++
		}
	}
	return sum
}

func part1(numbers []Number) int {
	move := []Number{}
	endIndex := (len(numbers) - 1) / 2
	for i := 0; i < len(numbers); i++ {
		if i > endIndex*2 {
			break
		}

		if i%2 == 0 {
			move = append(move, numbers[i])
		} else {
			left := numbers[i]
			for {
				if endIndex*2 < i {
					break
				}
				right := numbers[endIndex*2]
				if right.Len == left.Len {
					move = append(move, right)

					numbers[endIndex*2].Len -= left.Len
					numbers[i].Len = 0
					endIndex--
					break
				} else if right.Len > numbers[i].Len {
					move = append(move, Number{right.Value, left.Len})

					numbers[endIndex*2].Len -= left.Len
					numbers[i].Len = 0
					break
				} else {
					move = append(move, right)
					left.Len -= right.Len

					numbers[endIndex*2].Len = 0
					numbers[i].Len = left.Len
					endIndex--
				}
			}
		}
	}

	return sum(move)
}

func appendMultiple(numbers [][]Number) []Number {
	result := []Number{}
	for _, n := range numbers {
		result = append(result, n...)
	}
	return result
}

func part2(numbers []Number) int {
	maxMoved := math.MaxInt
	for i := (len(numbers) - 1); i > 1; {
		end := numbers[i]
		if end.Value > maxMoved || end.Value == -1 {
			i--
			continue
		}
		for n := 1; n < len(numbers); n++ {
			if numbers[n].Value == -1 {
				if numbers[n].Len == end.Len {
					maxMoved = end.Value
					numbers = appendMultiple([][]Number{
						numbers[:n],
						{end},
						numbers[n+1 : i],
						{{-1, end.Len}},
						numbers[i+1:],
					})
					print(numbers)
					i -= 1
					break
				} else if numbers[n].Len > end.Len {
					maxMoved = end.Value
					numbers = appendMultiple([][]Number{
						numbers[:n],
						{end},
						{{-1, numbers[n].Len - end.Len}},
						numbers[n+1 : i],
						{{-1, end.Len}},
						numbers[i+1:],
					})
					print(numbers)
					i -= 1
					break
				}
			}
			if n >= i {
				maxMoved = end.Value
				i--
				break
			}
		}
	}

	return sum(numbers)
}

func solution() (int, int) {
	numbers := getInput()
	numbers2 := make([]Number, len(numbers))
	copy(numbers2, numbers)
	p1 := part1(numbers)
	p2 := part2(numbers2)

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
