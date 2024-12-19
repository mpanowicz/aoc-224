package main

import (
	"aoc/internal/helpers"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Onsen struct {
	Stripes       []string
	StripesPrefix map[byte][]string

	Towels []string
}

func create(stripes []string, towels []string) Onsen {
	prefix := map[byte][]string{}
	for _, stripe := range stripes {
		prefix[stripe[0]] = append(prefix[stripe[0]], stripe)
	}
	for _, p := range prefix {
		sort.Slice(p, func(i, j int) bool {
			return len(p[i]) > len(p[j])
		})
	}
	return Onsen{stripes, prefix, towels}
}

func getInput() Onsen {
	f, _ := os.Open("cmd/day19/input.txt")
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	stripes := strings.Split(scanner.Text(), ", ")
	scanner.Scan()
	towels := []string{}
	for scanner.Scan() {
		towels = append(towels, scanner.Text())
	}
	return create(stripes, towels)
}

func (o *Onsen) CheckTowelPrefix(towel string, prefix []string, invalid *map[string]struct{}, valid *map[string]int) int {
	count := 0
	for _, p := range prefix {
		if strings.HasPrefix(towel, p) {
			part := towel[len(p):]
			if _, found := (*invalid)[part]; found {
				continue
			}
			if _, found := (*valid)[part]; found {
				count += (*valid)[part]
				continue
			}
			if len(part) == 0 {
				count++
			} else {
				count += o.CheckTowel(part, invalid, valid)
			}
		}
	}
	if count == 0 {
		(*invalid)[towel] = struct{}{}
	} else {
		(*valid)[towel] = count
	}
	return count
}

func (o *Onsen) CheckTowel(towel string, invalid *map[string]struct{}, valid *map[string]int) int {
	if s, found := o.StripesPrefix[towel[0]]; found {
		return o.CheckTowelPrefix(towel, s, invalid, valid)
	}

	return 0
}

func (o Onsen) CheckTowels() (int, int) {
	p1 := 0
	p2 := 0
	invalid := map[string]struct{}{}
	valid := map[string]int{}
	for _, towel := range o.Towels {
		check := o.CheckTowel(towel, &invalid, &valid)
		if check > 0 {
			fmt.Printf("OK: %s\n", towel)
			p1++
			fmt.Println(check)
			p2 += check
		} else {
			fmt.Printf("Impossible: %s\n", towel)
		}
	}
	return p1, p2
}

func solution() (int, int) {
	onsen := getInput()
	p1, p2 := onsen.CheckTowels()

	return p1, p2
}

func main() {
	helpers.PrintResult(solution())
}
