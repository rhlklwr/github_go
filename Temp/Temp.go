package main

import (
	"fmt"
	"os"
)

func main() {
	var v string

	for _, v = range os.Args[1:] {
		fmt.Printf("%q\n", v)
	}
	fmt.Println(v)
}
