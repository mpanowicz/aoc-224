package helpers

import "fmt"

func PrintResult(solution ...any) {
	p1, p2 := solution[0], solution[1]

	fmt.Printf("Part 1 result: %v\n", p1)
	fmt.Printf("Part 2 result: %v\n", p2)
}
