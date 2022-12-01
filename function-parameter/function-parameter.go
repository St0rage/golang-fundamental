package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Dani"
	lastName := "Yudistira"

	sayHelloTo(firstName, lastName)
	sayHelloTo("Dian", "Nugraha")
}