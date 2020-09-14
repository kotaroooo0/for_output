package main

import (
	"fmt"

	"github.com/golang/go/src/strings"
)

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	for k, v := range m {
		fmt.Print(k, v, " ")
	}
	strings.Join()
}
