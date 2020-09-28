package main

import "fmt"

func add(a int, b int) int {
	if b == 0 {
		return 0
	} else {
		return a + add(a, b-1)
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func helper(a int, b int, cur int) int {
	if b == 0 {
		return cur
	}
	return helper(a, b-1, a+cur)
}

func addIter(a, b int) int {
	return helper(a, b, 0)
}

func main() {
	c := addIter(4, 10)
	fmt.Println(c)

	d := fib(6)
	fmt.Println(d)
}
