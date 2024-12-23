package main

import (
	"aoc/internal/helpers"
	"bufio"
	"os"
	"sort"
	"strings"
)

type Connection struct {
	From string
	To   string
}

type Connections map[string][]string

type ConnectionsMap map[Connection]bool

type Input struct {
	Connections    Connections
	ConnectionsMap ConnectionsMap
}

func getInput() Input {
	f, _ := os.Open("cmd/day23/input.txt")
	scanner := bufio.NewScanner(f)

	connections := Connections{}
	connectionsMap := ConnectionsMap{}

	for scanner.Scan() {
		l := scanner.Text()
		parts := strings.Split(l, "-")
		if c, ok := connections[parts[0]]; ok {
			connections[parts[0]] = append(c, parts[1])
		} else {
			connections[parts[0]] = []string{parts[1]}
		}
		if c, ok := connections[parts[1]]; ok {
			connections[parts[1]] = append(c, parts[0])
		} else {
			connections[parts[1]] = []string{parts[0]}
		}

		connectionsMap[Connection{parts[0], parts[1]}] = true
		connectionsMap[Connection{parts[1], parts[0]}] = true
	}

	return Input{connections, connectionsMap}
}

func sortAndConcat(c []string) string {
	sort.Slice(c, func(i, j int) bool {
		return c[i] < c[j]
	})
	return strings.Join(c, ",")
}

func contains(c []string, s string) bool {
	for _, v := range c {
		if v == s {
			return true
		}
	}
	return false
}

func countLANofThree(conns Connections) int {
	visited := map[string]bool{}

	for from, vals := range conns {
		for _, v := range vals {
			for _, vVal := range conns[v] {
				if vVal != from {
					vValVal := conns[vVal]
					if contains(vValVal, from) && (from[0] == 't' || v[0] == 't' || vVal[0] == 't') {
						con := []string{from, v, vVal}
						conStr := sortAndConcat(con)
						if _, ok := visited[conStr]; !ok {
							visited[conStr] = true
						}
					}
				}
			}
		}
	}

	return len(visited)
}

func (i Input) intersection(x, y []string) []string {
	xMap := map[string]struct{}{}
	for _, v := range x {
		xMap[v] = struct{}{}
	}
	in := []string{}
	for _, v := range y {
		if _, ok := xMap[v]; ok {
			in = append(in, v)
		}
	}
	// fmt.Println("x", x)
	// fmt.Println("y", y)
	// fmt.Println("i", in)

	return in
}

func startsWithT(s []string) bool {
	for _, v := range s {
		if v[0] == 't' {
			return true
		}
	}
	return false
}

func (in Input) bronKerbosch(r, p, x []string) [][]string {
	cliques := [][]string{}
	if len(p) == 0 && len(x) == 0 {
		cliques = append(cliques, r)
		return cliques
	}

	for len(p) > 0 {
		v := p[0]
		newR := make([]string, len(r)+1)
		copy(newR, r)
		newR[len(r)] = v
		bk := in.bronKerbosch(newR, in.intersection(p, in.Connections[v]), in.intersection(x, in.Connections[v]))
		cliques = append(cliques, bk...)
		p = p[1:]
		x = append(x, v)
	}

	return cliques
}

func (i Input) getClique() []string {
	nodes := []string{}
	for k := range i.Connections {
		nodes = append(nodes, k)
	}
	bk := i.bronKerbosch([]string{}, nodes, []string{})
	sort.Slice(bk, func(i, j int) bool {
		return len(bk[i]) > len(bk[j])
	})
	var p2 []string
	for _, v := range bk {
		if startsWithT(v) {
			p2 = v
			break
		}
	}
	sort.Slice(p2, func(i, j int) bool {
		return p2[i] < p2[j]
	})

	return p2
}

func solution() (int, string) {
	input := getInput()
	p1 := countLANofThree(input.Connections)
	cycle := input.getClique()

	return p1, strings.Join(cycle, ",")
}

func main() {
	helpers.PrintResult(solution())
}
