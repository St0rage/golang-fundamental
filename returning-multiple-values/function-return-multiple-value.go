package main

import "fmt"

func getFullName() (string, string, string) {
	return "Dani", "Yudistira", "Maulana"
}

func main() {
	// firstName, middleName, lastName := getFullName()
	// fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)

	// fmt.Println(getFullName())

	// gunakan "_" (underscore) untuk ignore variable
	firstName, _, _ := getFullName()
	fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)

	fmt.Println(getFullName())

	
}