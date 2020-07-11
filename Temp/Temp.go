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
	blue := []int{
		4, 5, 6, 7, 8, 10,
	}
	blue = append(blue, 8, 9, 10, 11, 12)
	fmt.Printf("blue: %#v\n%v\n%v", blue, cap(blue), len(blue))
}
