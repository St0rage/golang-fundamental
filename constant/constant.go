package main

import "fmt"

func main() {
	// const firtName string = "Dani"
	// const lastName = "Yudistira"
	// const value = 1000

	// fmt.Println(firtName)	
	// fmt.Println(lastName)	
	// fmt.Println(value)

	// multiple constant
	const (
		firstName = "Dani"
		lastName = "Yudistira" 
		value = 1000
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}