package main

import "fmt"

func main() {
	name := "Dani"
	counter := 0

	increment := func() {
		name := "Dian"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}