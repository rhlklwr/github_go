package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string

	dir, file = path.Split("Path/main.css")

	fmt.Println("dir : ", dir)
	fmt.Println("file : ", file)
}
