package main

import (
	"fmt"
	"strconv"
)

func add(x int, y int) int {
	return x + y
}

func add2(x, y, z int, d string) string {
	return d + strconv.Itoa(x+y+z)
}

func subprod(x, y int) (int, int) {
	return x - y, x * y
}

func many(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}

func main() {
	c := add(4, 5)
	fmt.Println(c)

	var f string = add2(4, 5, 6, "sum: ")
	fmt.Println(f)

	g, h := subprod(5, 4)
	fmt.Println(g, h)

	j := many(1, 2, 3, 4, 5, 6)
	fmt.Println(j)
}
