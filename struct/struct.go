package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var dani Customer
	dani.Name = "Dani"
	dani.Address = "Indonesia"
	dani.Age = 30

	fmt.Println(dani)
	fmt.Println(dani.Name)
	fmt.Println(dani.Address)
	fmt.Println(dani.Age)

	// Struct Literals
	dian := Customer{
		Name: "Dian",
		Address: "Garut",
		Age: 5,
	}
	fmt.Println(dian)

	// cara ini tidak direkomendasikan
	melani := Customer{"Melani", "Garut", 13}
	fmt.Println(melani)
}