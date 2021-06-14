package main

import (
	"fmt"
)
var y = 42
var z string = "shaken, not stirred"
var a string = 'james said"shaken, in stirred"'

func main() {



	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(a)
	fmt.Printf("%T\n", a)
}