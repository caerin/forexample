package main

import (
	"fmt"
)

func main() {
	// init ; condition; post {}
	x := 0
	for {
		//x++
		if x > 100 {

			break
		}
		if x%2 != 0 {
			continue
		}

		fmt.Println(x)
		x++
	}
	fmt.Println("DONE!")
}
