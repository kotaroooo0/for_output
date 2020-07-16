package main

import (
	"fmt"
	"reflect"
)

func main() {
	//validate(false)
	//refelcttest()

	p := &Point{5, 10}
	m := toMap(p)

	pp := &Point{}
	toStruct(m, pp)
	fmt.Println(pp)

	u := &User{123, "kotaroooo0", 24}
	mm := toMap(u)
	fmt.Println(mm)

	uu := &User{}
	toStruct(mm, uu)
	fmt.Println(uu)

}

type User struct {
	Id   int
	Name string
	Age  int
}

func toMap(v interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	rv := reflect.ValueOf(v).Elem()
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		ftv := rt.Field(i)
		fv := rv.Field(i)
		if rv.CanSet() {
			m[ftv.Name] = fv.Interface()
		}
	}
	return m
}

func toStruct(m map[string]interface{}, s interface{}) {
	rv := reflect.ValueOf(s).Elem()
	for k, v := range m {
		rv.FieldByName(k).Set(reflect.ValueOf(v))
	}
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
