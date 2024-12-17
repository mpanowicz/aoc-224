package main

import (
	"aoc/internal/helpers"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Opcode int
type Operand int
type Instruction struct {
	Opcode  Opcode
	Operand Operand
}

var (
	adv Opcode = 0
	bxl Opcode = 1
	bst Opcode = 2
	jnz Opcode = 3
	bxc Opcode = 4
	out Opcode = 5
	bdv Opcode = 6
	cdv Opcode = 7

	a Operand = 4
	b Operand = 5
	c Operand = 6

	debugEnabled bool = false
)

type Computer struct {
	A int
	B int
	C int

	Instructions []Instruction
	Pointer      int
	Out          []int
}

func getInput() Computer {
	f, _ := os.Open("cmd/day17/input.txt")
	scanner := bufio.NewScanner(f)

	c := Computer{}
	scanner.Scan()
	c.A = helpers.ParseInt(strings.Split(scanner.Text(), ": ")[1])
	scanner.Scan()
	c.B = helpers.ParseInt(strings.Split(scanner.Text(), ": ")[1])
	scanner.Scan()
	c.C = helpers.ParseInt(strings.Split(scanner.Text(), ": ")[1])

	scanner.Scan()
	scanner.Scan()
	l := strings.Split(scanner.Text(), " ")
	parts := strings.Split(l[1], ",")
	c.Instructions = make([]Instruction, len(parts)/2)
	for i := 0; i < len(parts); i += 2 {
		opcode := parts[i]
		operand := parts[i+1]
		c.Instructions[i/2] = Instruction{
			Opcode:  Opcode(helpers.ParseInt(opcode)),
			Operand: Operand(helpers.ParseInt(operand)),
		}
	}

	return c
}

func (c *Computer) debug() {
	if debugEnabled {
		fmt.Println(c.A, c.B, c.C, c.Pointer, c.Output())
	}
}

func (c *Computer) execute(check []int) {
	c.Out = []int{}

	checkEnabled := len(check) > 0

	c.debug()
	for c.Pointer < len(c.Instructions) {
		instruction := c.Instructions[c.Pointer]
		operandValue := instruction.Operand.Value(c)

		switch instruction.Opcode {
		case adv:
			c.A = c.A / int(math.Pow(2, float64(operandValue)))
		case bdv:
			c.B = c.A / int(math.Pow(2, float64(operandValue)))
		case cdv:
			c.C = c.A / int(math.Pow(2, float64(operandValue)))
		case bxl:
			c.B = c.B ^ int(instruction.Operand)
		case bst:
			c.B = operandValue % 8
		case jnz:
			if c.A != 0 {
				c.Pointer = operandValue
				c.debug()
				continue
			}
		case bxc:
			c.B = c.B ^ c.C
		case out:
			c.Out = append(c.Out, operandValue%8)

			if checkEnabled {
				if len(c.Out) > len(check) {
					return
				}
				for i, v := range c.Out {
					if v != check[i] {
						return
					}
				}
			}
		}
		c.debug()
		c.Pointer++
	}
}

func (o Operand) Value(comp *Computer) int {
	if o == a {
		return comp.A
	} else if o == b {
		return comp.B
	} else if o == c {
		return comp.C
	}

	return int(o)
}

func (c *Computer) InstructionInts() []int {
	ints := []int{}
	for _, i := range c.Instructions {
		ints = append(ints, int(i.Opcode), int(i.Operand))
	}
	return ints
}

func (c *Computer) Output() string {
	s := []string{}
	for _, o := range c.Out {
		s = append(s, fmt.Sprint(o))
	}
	return strings.Join(s, ",")
}

func step(a int) int {
	b := a % 8
	b = b ^ 6
	c := a / int(math.Pow(2, float64(b)))
	b = b ^ c
	b = b ^ 4
	return b % 8
}

func findCopy(ins []int, a, n int) int {
	if step(a) != ins[len(ins)-1-n] {
		return -1
	}

	if n == len(ins)-1 {
		return a
	} else {
		for i := range 8 {
			x := findCopy(ins, a*8+i, n+1)
			if x != -1 {
				return x
			}
		}
	}

	return -1
}

func find(ins []int) int {
	for i := range 8 {
		x := findCopy(ins, i, 0)
		if x != -1 {
			return x
		}
	}
	return -1
}

func solution() (string, int) {
	c := getInput()
	c.execute([]int{})
	p2 := find(c.InstructionInts())

	return c.Output(), p2
}

func main() {
	helpers.PrintResult(solution())
}
