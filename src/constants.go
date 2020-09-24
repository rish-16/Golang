package main

import (
	"fmt"
	"math"
)

const s string = "my_constant"

func main() {
	fmt.Println(s)

	const n = 500000000
	fmt.Println(n)

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(math.Sin(d))
}
