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
