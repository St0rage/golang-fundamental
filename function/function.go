package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World!")
}

func main() {
	sayHello()
	sayHello()
	sayHello()
	fmt.Printf("\n\n")

	for i := 1; i <= 10; i++ {
		sayHello()
	}
}