package main

import (
	"fmt"
)

var y = 42

type hotdog int

var b hotdog
var z string = "shaken, not stirred"

var a string = `Hello,"guys"`

func main() {

	b = 43

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	y = int(b)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
