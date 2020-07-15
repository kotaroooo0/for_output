package main

import (
	"fmt"
	"reflect"
)

func main() {
	validate(false)
	refelcttest()
}

func validate(x interface{}) error {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")

	case map[string]int:
		fmt.Println("map")
	default:
		fmt.Println("golang")
	}

	s, ok := x.(string)
	if ok {
		fmt.Println(s)
	}
	return nil
}

type Point struct {
	X int
	Y int
}

func refelcttest() {
	p := &Point{X: 10, Y: 5}
	rv := reflect.ValueOf(p).Elem()
	fmt.Println(rv.Type())
	fmt.Println(rv.Kind())
	fmt.Println(rv.Interface())

	xv := rv.Field(0)
	fmt.Println(xv.Int())
	xv.SetInt(100)
	fmt.Println(xv.Int())
}
