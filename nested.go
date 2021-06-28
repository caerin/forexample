package main

import (
	"fmt"
)

func main() {
	// init ; condition; post {}
	for i := 0; i <= 100; i++ {

		fmt.Printf("the outer loop: %d\n", i)

		for j := 0; j < 3; j++ {

			fmt.Printf("\t\tThe inner loop: %d\n", j)
		}
		//fmt.Println(i)
	}

}
