package main

import "fmt"

type printer interface {
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

	for _, s := range z {
		switch e := s.(type) {
		case book:
			e.print()
		case game:
			e.print()
			e.discount()
		}

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
