package main

import "fmt"

func main() {

	c := make(chan int, 5)
	cr := make(chan<- int, 5)
	cs := make(<-chan int, 5)

	_ = cr
	_ = cs
	c <- 45

	fmt.Println(c)

}
