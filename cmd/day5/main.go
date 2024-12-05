package main

import (
	"aoc/internal/helpers"
	"bufio"
	"os"
	"strings"
)

type PageOrdering struct {
	Ordering map[int]map[int]struct{}
	Updates  [][]int
}

func getInput() PageOrdering {
	f, _ := os.Open("cmd/day5/input.txt")
	s := bufio.NewScanner(f)

	orderingRead := false
	ordering := make(map[int]map[int]struct{})
	updates := [][]int{}
	for {
		if !s.Scan() {
			break
		}
		line := s.Text()
		if len(line) == 0 {
			orderingRead = true
			continue
		}

		if !orderingRead {
			parts := strings.Split(line, "|")
			before := helpers.ParseInt(parts[0])
			after := helpers.ParseInt(parts[1])

			if _, ok := ordering[before]; ok {
				ordering[before][after] = struct{}{}
			} else {
				ordering[before] = map[int]struct{}{after: {}}
			}
		} else {
			parts := strings.Split(line, ",")
			pages := make([]int, len(parts))
			for i, p := range parts {
				pages[i] = helpers.ParseInt(p)
			}
			updates = append(updates, pages)
		}
	}

	return PageOrdering{ordering, updates}
}

func (po PageOrdering) CheckPageOrdering(page int, before []int) bool {
	if after, ok := po.Ordering[page]; ok {
		for _, v := range before {
			if _, ok := after[v]; ok {
				return false
			}
		}
	}
	return true
}

func (po PageOrdering) CheckUpdateOrdering(update []int) bool {
	for i := len(update) - 1; i >= 0; i-- {
		page := update[i]
		if !po.CheckPageOrdering(page, update[:i]) {
			return false
		}
	}
	return true
}

func (po PageOrdering) Sort(update []int) []int {
	i := len(update) - 1
	for i >= 0 {
		page := update[i]
		if !po.CheckPageOrdering(page, update[:i]) {
			break
		}
		i--
	}
	updated := update[i+1:]
	for u := i; u >= 0; u-- {
		temp := make([]int, len(updated)+1)
		for i := len(updated); i >= 0; i-- {
			copy(temp, updated[:i])
			temp[i] = update[u]
			copy(temp[i+1:], updated[i:])
			if po.CheckUpdateOrdering(temp) {
				updated = temp
				break
			}
		}
	}
	return updated
}

func solution() (int, int) {
	po := getInput()

	validSum := 0
	invalidUpdates := [][]int{}
	for _, update := range po.Updates {
		if po.CheckUpdateOrdering(update) {
			middle := int(float64(len(update) / 2))
			validSum += update[middle]
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	correctedSum := 0
	for _, update := range invalidUpdates {
		sorted := po.Sort(update)
		middle := int(float64(len(sorted) / 2))
		correctedSum += sorted[middle]
	}

	return validSum, correctedSum
}

func main() {
	helpers.PrintResult(solution())
}
