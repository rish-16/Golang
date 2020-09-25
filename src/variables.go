package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = "hello"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b + c)

	var d bool = false
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d))

	var e int
	fmt.Println(e)

	f := "world"
	println(a + f)
}
