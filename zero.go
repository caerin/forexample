package main

import (
	"fmt"
)

var y string
var z int

func main() {

	fmt.Println("printing the value of y:", y, "ending")
	fmt.Printf("%T\n", y)

	y = "Bond, James BOnd"
	fmt.Println("printing the value of y:", y)
	fmt.Printf("%T\n", y)
	fmt.Println("value of z:", z)
	fmt.Printf("%T\n", z)
}
