package main

import (
	"fmt"
)

func main() {
	// init ; condition; post {}
	x := 1
	for x < 10 {

		fmt.Println(x)
		x++

	}
	fmt.Println("DONE!")
}
