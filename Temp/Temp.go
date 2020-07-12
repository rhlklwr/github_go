package main

import (
	"fmt"
)

func main() {

	var hey = []int{1, 2, 3}
	var bye = []int{2, 3}

	var dict = map[string][]int{
		"Hello": hey,
	}

	dict["bye"] = bye

	turkish := make(map[string][]int, len(dict))

	var turkish1 map[string][]int

	if turkish == nil {
		fmt.Println("turkish")
	}

	if turkish1 == nil {
		fmt.Println("turkish1")
	}

}
