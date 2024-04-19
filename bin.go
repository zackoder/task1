package moudel

import (
	"fmt"
	"strconv"
)

func Bin(s string) string {
	a, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println("Error", err)
	}
	return strconv.Itoa(int(a))
}