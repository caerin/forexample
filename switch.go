package main

import (
	"fmt"
)

func main() {
	n := "Bond"
	switch n {

	case "James", "Bond", "Money":

		fmt.Println("James or Bond or Money")

	case "M":

		fmt.Println("this print")
		//fallthrough
	case "N":

		fmt.Println("they are equal")
		//fallthrough
	default:

		fmt.Println("this is default")
	}
}
