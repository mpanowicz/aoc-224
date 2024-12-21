package main

import (
	"aoc/internal/helpers"
	"bufio"
	"os"
)

type Point struct {
	X, Y int
}

type Button struct {
	Position Point
	Value    rune
}

type Keypad struct {
	Buttons map[rune]Point
}

var (
	Numeric = Keypad{
		Buttons: map[rune]Point{
			'0': {1, 0},
			'A': {2, 0},
			'1': {0, 1},
			'2': {1, 1},
			'3': {2, 1},
			'4': {0, 2},
			'5': {1, 2},
			'6': {2, 2},
			'7': {0, 3},
			'8': {1, 3},
			'9': {2, 3},
		},
	}

	Directional = Keypad{
		Buttons: map[rune]Point{
			'<': {0, 0},
			'v': {1, 0},
			'>': {2, 0},
			'^': {1, 1},
			'A': {2, 1},
		},
	}
)

func getInput() []string {
	numbers := []string{}

	f, _ := os.Open("cmd/day21/input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}

	return numbers
}

func getNumericPresses(number string) string {
	current := Numeric.Buttons['A']
	output := []rune{}
	for _, r := range number {
		next := Numeric.Buttons[r]
		difX, difY := next.X-current.X, next.Y-current.Y

		horizontal := []rune{}
		for i := 0; i < helpers.Abs(difX); i++ {
			if difX >= 0 {
				horizontal = append(horizontal, '>')
			} else {
				horizontal = append(horizontal, '<')
			}
		}

		vertical := []rune{}
		for i := 0; i < helpers.Abs(difY); i++ {
			if difY >= 0 {
				vertical = append(vertical, '^')
			} else {
				vertical = append(vertical, 'v')
			}
		}

		if current.Y == 0 && next.X == 0 {
			output = append(output, vertical...)
			output = append(output, horizontal...)
		} else if current.X == 0 && next.Y == 0 {
			output = append(output, horizontal...)
			output = append(output, vertical...)
		} else if difX < 0 {
			output = append(output, horizontal...)
			output = append(output, vertical...)
		} else {
			output = append(output, vertical...)
			output = append(output, horizontal...)
		}

		current = next
		output = append(output, 'A')
	}

	return string(output)
}

func getDirectionalPresses(number string) string {
	current := Directional.Buttons['A']
	output := []rune{}

	for _, r := range number {
		next := Directional.Buttons[r]
		difX, difY := next.X-current.X, next.Y-current.Y

		horizontal := []rune{}
		for i := 0; i < helpers.Abs(difX); i++ {
			if difX >= 0 {
				horizontal = append(horizontal, '>')
			} else {
				horizontal = append(horizontal, '<')
			}
		}

		vertical := []rune{}
		for i := 0; i < helpers.Abs(difY); i++ {
			if difY >= 0 {
				vertical = append(vertical, '^')
			} else {
				vertical = append(vertical, 'v')
			}
		}

		if current.X == 0 && next.Y == 1 {
			output = append(output, horizontal...)
			output = append(output, vertical...)
		} else if current.Y == 1 && next.X == 0 {
			output = append(output, vertical...)
			output = append(output, horizontal...)
		} else if difX < 0 {
			output = append(output, horizontal...)
			output = append(output, vertical...)
		} else {
			output = append(output, vertical...)
			output = append(output, horizontal...)
		}
		current = next
		output = append(output, 'A')
	}

	return string(output)
}

func MultipleRobots(number string, maxRobots int, robot int, cache map[string][]int) int {
	if val, ok := cache[number]; ok {
		if val[robot-1] != 0 {
			return val[robot-1]
		}
	} else {
		cache[number] = make([]int, maxRobots)
	}

	directional := getDirectionalPresses(number)
	cache[number][0] = len(directional)

	if robot == maxRobots {
		return len(directional)
	}

	parts := splitParts(directional)

	count := 0
	for _, part := range parts {
		c := MultipleRobots(part, maxRobots, robot+1, cache)
		if _, ok := cache[part]; !ok {
			cache[part] = make([]int, maxRobots)
		}
		cache[part][0] = c
		count += c
	}

	cache[number][robot-1] = count
	return count
}

func splitParts(number string) []string {
	parts := []string{}
	current := []rune{}
	for _, r := range number {
		current = append(current, r)

		if r == 'A' {
			parts = append(parts, string(current))
			current = []rune{}
		}
	}
	return parts
}

func sequence(nums []string, robots int, cache map[string][]int) int {
	count := 0
	for _, num := range nums {
		direction := getNumericPresses(num)
		val := MultipleRobots(direction, robots, 1, cache)
		numeric := helpers.ParseInt(num[:len(num)-1])
		count += numeric * val
	}
	return count
}

func solution() (int, int) {
	nums := getInput()

	cache := map[string][]int{}
	p1 := sequence(nums, 2, cache)
	cache = map[string][]int{}
	p2 := sequence(nums, 25, cache)

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
