package main

import (
	"fmt"
)

/*
const (
	a = 5
	b = "Hello"
)
*/
var (
	a = 5
	b = "Hello"
)

func Main() {

	c := "Detect if it is string"
	fmt.Printf("Type is : %T\n", c)
	fmt.Println(a, b, c)
}
