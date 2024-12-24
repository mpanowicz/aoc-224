package main

import (
	"aoc/internal/helpers"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Gate string

var (
	And = Gate("AND")
	Or  = Gate("OR")
	Xor = Gate("XOR")
)

type Connection struct {
	Left   string
	Right  string
	Output string
	Gate   Gate
}

type Input struct {
	Connections []Connection
	WireValues  map[string]int
}

type ZWire struct {
	Idx   int
	Value int
}

func getInput() Input {
	f, _ := os.Open("cmd/day24/input.txt")
	scanner := bufio.NewScanner(f)

	i := Input{
		[]Connection{},
		map[string]int{},
	}
	readWires := true
	for scanner.Scan() {
		if scanner.Text() == "" {
			readWires = false
			continue
		}

		if readWires {
			l := scanner.Text()
			parts := strings.Split(l, ": ")
			i.WireValues[parts[0]] = helpers.ParseInt(parts[1])
		} else {
			l := scanner.Text()
			fromTo := strings.Split(l, " -> ")
			output := fromTo[1]
			parts := strings.Split(fromTo[0], " ")
			left := parts[0]
			right := parts[2]
			var gate Gate
			switch parts[1] {
			case "AND":
				gate = And
			case "OR":
				gate = Or
			case "XOR":
				gate = Xor
			}
			i.Connections = append(i.Connections, Connection{left, right, output, gate})
		}
	}
	return i
}

func (g Gate) result(l, r int) int {
	switch g {
	case And:
		return l & r
	case Or:
		return l | r
	case Xor:
		return l ^ r
	}
	return -1
}

func (i *Input) findValues() {
	conns := i.Connections
	for len(conns) > 0 {
		tempConns := []Connection{}
		for _, c := range conns {
			l, lOk := i.WireValues[c.Left]
			r, rOk := i.WireValues[c.Right]
			if !lOk || !rOk {
				tempConns = append(tempConns, c)
				continue
			}

			i.WireValues[c.Output] = c.Gate.result(l, r)
		}
		conns = tempConns
	}
}

func (i *Input) getZ() int {
	zs := []ZWire{}
	for k, v := range i.WireValues {
		if k[0] == 'z' {
			i := helpers.ParseInt(k[1:])
			zs = append(zs, ZWire{i, v})
		}
	}
	sort.Slice(zs, func(i, j int) bool {
		return zs[i].Idx > zs[j].Idx
	})
	values := ""
	for _, w := range zs {
		values += fmt.Sprintf("%d", w.Value)
	}
	fmt.Println(values)
	v, _ := strconv.ParseInt(values, 2, 0)
	return int(v)
}

func (i *Input) findInvalidZ() []string {
	zs := []string{}
	for _, w := range i.Connections {
		if w.Output[0] == 'z' && w.Gate != Xor && w.Output != "z45" {
			zs = append(zs, w.Output)
		}
	}
	return zs
}

func notXY(s string) bool {
	return s[0] != 'x' && s[0] != 'y'
}

func (i *Input) findXorOrIns() map[string]struct{} {
	xors := map[string]struct{}{}
	for _, v := range i.Connections {
		if v.Gate == Xor {
			xors[v.Output] = struct{}{}
		}
	}
	return xors
}

func (i *Input) findAndOrIns() map[string]struct{} {
	ands := map[string]struct{}{}
	for _, v := range i.Connections {
		if v.Gate == And && v.Left != "x00" && v.Right != "x00" {
			ands[v.Output] = struct{}{}
		}
	}
	return ands
}

func (i *Input) findInvalidOrs() []string {
	notInOrs := i.findXorOrIns()
	inOrs := i.findAndOrIns()
	ors := []string{}
	for _, c := range i.Connections {
		if c.Gate == Or {
			_, lOk := notInOrs[c.Left]
			if lOk {
				ors = append(ors, c.Left)
			}
			_, rOk := notInOrs[c.Right]
			if rOk {
				ors = append(ors, c.Right)
			}
		} else {
			_, lOk := inOrs[c.Left]
			if lOk {
				ors = append(ors, c.Left)
			}
			_, rOk := inOrs[c.Right]
			if rOk {
				ors = append(ors, c.Right)
			}
		}
	}
	return ors
}

func (i *Input) findInvalidXors() []string {
	xors := []string{}
	for _, w := range i.Connections {
		if w.Gate == Xor {
			if notXY(w.Left) && notXY(w.Right) && w.Output[0] != 'z' {
				xors = append(xors, w.Output)
			}
		}
	}
	return xors
}

func (i *Input) findInvalid() string {
	is := map[string]struct{}{}

	iz := i.findInvalidZ()
	for _, c := range iz {
		is[c] = struct{}{}
	}
	ix := i.findInvalidXors()
	for _, c := range ix {
		is[c] = struct{}{}
	}
	io := i.findInvalidOrs()
	for _, c := range io {
		is[c] = struct{}{}
	}

	keys := []string{}
	for k := range is {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return strings.Join(keys, ",")
}

func solution() (int, string) {
	i := getInput()
	i.findValues()
	p1 := i.getZ()
	p2 := i.findInvalid()

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
