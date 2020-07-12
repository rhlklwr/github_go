package main

import "fmt"

func main() {

	var a = map[string]int{"A": 1}
	fmt.Println(a)
	var b = map[string]int{"A": 5}

	c := a

	c["A"] = 6

	fmt.Println(a, b, c)

}
