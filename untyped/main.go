package main

import (
	"fmt"
)

func main() {
	x := 1
	y := 0.1
	fmt.Println(x + y)
	// invalid operation: x + y (mismatched types int and float64)

	fmt.Println(1 + 0.1) // 1.1
	fmt.Println(1 + y)   // 1.1

	fmt.Println(x + 0.1)
	// constant 0.1 truncated to integer

	type MyInt int
	var myInt MyInt = 42
	fmt.Println(10 + myInt) // 52

	fmt.Println(x + myInt)
	// invalid operation: mismatched types int and MyInt
}
