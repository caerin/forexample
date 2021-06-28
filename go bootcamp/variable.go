package main

import (
	"fmt"
)

var x = 42
var y = 43
var z = 123

func main() {
	//declare a variable and assign a value

	//x := 42
	fmt.Println(x)
	//x = 99

	fmt.Println(y)
	foo()
	fmt.Println(z)

}
func foo() {

	fmt.Println("again:", y)
}
