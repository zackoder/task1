package main

import (
	"fmt"
)

func main() {
	
	s := "" 
	runes := []rune(s)
	r := ""
	for i := 0; i < len(s); i++ {
		if runes[i] == ' '  && runes[i+1] == '\'' {
			r += string(runes[i+1])
			i++
		} else if runes[i] == ' ' && runes[i-1] == '\'' {
			r += string(runes[i+1])
			i++
		} else {
			r += string(runes[i])
		}
	}
	fmt.Println(r)
}