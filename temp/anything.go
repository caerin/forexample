package main

import "fmt"

func main() {
	fmt.Println("this is the most fun and learning the go programming language")
	foo()
	fmt.Println("something more")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("im in foo")
}
func bar() {
	fmt.Println("and we exited")
}
