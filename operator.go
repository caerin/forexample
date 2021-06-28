package main

import (
	"fmt"
)

func main() {

	fmt.Println("%v\n", true && true)

	fmt.Println("%v\n", true && false)
	fmt.Println("%v\n", true || true)
	fmt.Println("%v\n", true || false)
	fmt.Println("%v\n", !true)

}
