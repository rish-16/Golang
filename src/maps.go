package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 4
	m["k1"] = 8

	fmt.Println(m)

	_, present := m["k3"]
	fmt.Println(present)

	n := map[string]int{"k3": 4, "k5": 10}
	fmt.Println(n)
}
