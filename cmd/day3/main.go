package main

import (
	"aoc/internal/helpers"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getInput() string {
	f, _ := os.Open("cmd/day3/input.txt")
	r := bufio.NewReader(f)

	input := []rune{}
	for {
		c, _, err := r.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}
		input = append(input, c)
	}
	return string(input)
}

type Pair struct {
	X int
	Y int
}

type Instruction struct {
	Type string
	Pair Pair
}

func updateIndex(i *int, change int, max int) bool {
	*i = *i + change
	return *i < max
}

func parse(input string) <-chan Instruction {
	ch := make(chan Instruction)

	left := "mul("
	middle := ","
	right := ")"
	do := "do()"
	dont := "don't()"

	go func() {
		for i := 0; i < len(input)-len(left); i++ {
			if input[i:i+len(do)] == do {
				ch <- Instruction{Type: "Do"}
			}
			if input[i:i+len(dont)] == dont {
				ch <- Instruction{Type: "Dont"}
			}

			check := input[i : i+len(left)]
			if check == left {
				if !updateIndex(&i, len(left), len(input)) {
					break
				}
				parts := strings.Split(input[i:], ",")
				numStr := parts[0]
				num, err := strconv.Atoi(numStr)
				if err != nil {
					continue
				}

				if !updateIndex(&i, len(numStr), len(input)) {
					break
				}
				check = string(input[i])
				if check != middle {
					i = i - 1
					continue
				}
				if !updateIndex(&i, 1, len(input)) {
					break
				}

				pair := Pair{}
				pair.X = num

				parts = strings.Split(input[i:], ")")
				numStr = parts[0]
				num, err = strconv.Atoi(numStr)
				if err != nil {
					continue
				}
				if !updateIndex(&i, len(numStr), len(input)) {
					break
				}
				check = string(input[i])
				if check == right {
					pair.Y = num
					ch <- Instruction{"Pair", pair}
				}
			}
		}
		close(ch)
	}()

	return ch
}

func solution() (int, int) {
	input := getInput()
	sumP1 := 0
	sumP2 := 0
	do := true
	for p := range parse(input) {
		if p.Type == "Pair" {
			sumP1 += p.Pair.X * p.Pair.Y
			if do {
				sumP2 += p.Pair.X * p.Pair.Y
			}
		} else if p.Type == "Do" {
			do = true
		} else {
			do = false
		}
	}

	return sumP1, sumP2
}

func main() {
	helpers.PrintResult(solution())
}
