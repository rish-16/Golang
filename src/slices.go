package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "3"
	s[1] = "2"
	s[2] = "1"
	fmt.Println(s)

	s = append(s, "4", "5", "6")
	fmt.Println(s)

	fmt.Println(s[1:])
}
