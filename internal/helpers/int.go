package helpers

import (
	"fmt"
	"strconv"
)

func ParseInt(s string) int {
	v, e := strconv.Atoi(s)
	if e != nil {
		fmt.Println(e)
	}
	return v
}

func ParseInts(s []string) []int {
	var v []int
	for _, i := range s {
		v = append(v, ParseInt(i))
	}
	return v
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
