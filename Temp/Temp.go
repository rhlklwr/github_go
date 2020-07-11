package main

import (
	"fmt"
)

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	blue := [...][3]int{
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Printf("blue: %#v\n", blue)
}
