package main

import "fmt"

type printer interface {
	print()
	discount()
}

type book struct {
	name  string
	price int
}

type game struct {
	name  string
	price float64
}

func main() {

	var book1 = book{
		name:  "Moby",
		price: 5,
	}

	var game1 = game{
		name:  "GTA",
		price: 5.5,
	}

	type list []printer

	var z list

	z = append(z, book1, game1)

	game1.discount()

	for _, h := range z {
		h.print()
		h.discount()
	}
}

func (b book) print() {
	fmt.Printf("Book name is %v\nBook price is %v\n", b.name, b.price)
}

func (g game) print() {
	fmt.Printf("Game name is %v\nGame price is %v\n", g.name, g.price)
}

func (g game) discount() {
	fmt.Printf("Game name is %v\nGame price is %v\n", g.name, g.price-10)
}

func (b book) discount() {
	fmt.Printf("Game name is %v\nGame price is %v\n", b.name, b.price-10)
}
