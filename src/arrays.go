package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)

	fmt.Println(len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b[3])

	var twoD [2][3]int
	twoD[1][1] = 45
	for i := 0; i < 2; i++ {
		var header string = ""
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
			header += " " + strconv.Itoa(twoD[i][j]) + " "
		}
		fmt.Println(header)
	}

	fmt.Println(twoD)
}
