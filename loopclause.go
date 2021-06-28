package main

import (
	"fmt"
)

func main() {
	// init ; condition; post {}
	i := 0
	for {
		if i > 9 {
			break

		}
		fmt.Println(i)
		i++
	}
	fmt.Println("DONE!")
}
