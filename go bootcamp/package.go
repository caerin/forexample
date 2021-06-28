package main

import (
	"fmt" "os" "io"
)

var y = 42

func main() {

	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%T\t%b\t%#x\n", y, y, y)
	s := fmt.Sprintf("%T\t%b\t%#x\n", y, y, y)
	fmt.Println(s)
	y = 911
	fmt.Printf("%v\n", y)
}
