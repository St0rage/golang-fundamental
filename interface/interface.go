package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasname HasName) {
	fmt.Println("Hello", hasname.GetName())
}

// Struct Person
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// Struct Animal
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var dani Person
	dani.Name = "Dani"
	SayHello(dani)

	cat := Animal{"Push"}
	SayHello(cat)
}
