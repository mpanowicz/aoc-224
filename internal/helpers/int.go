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
