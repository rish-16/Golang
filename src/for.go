package main

import "fmt"

func main() {

	var i int = 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 9; j++ {
		fmt.Println(j)
	}

	// while loop
	for {
		fmt.Println("loop")
		break
	}

	// skipping indices
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
