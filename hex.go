package moudel

import (
	"fmt"
	"strconv"
)

 func Hex(s string) int64 {
	a , err:= strconv.ParseInt(s, 16,64)
	if err != nil {
		fmt.Println("Error", err)
}
 	return a
}


