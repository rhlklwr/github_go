package main

import "fmt"

const max = 20

func main() {
	fmt.Printf("%5s", "X")
	for i := 1; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 1; i <= max; i++ {
		fmt.Printf("%5d", i)

		// print the cells
		for j := 1; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
