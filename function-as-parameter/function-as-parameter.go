package main

import (
	"fmt"
)

type Filter func(string)string //type declaration function

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string)string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
 
func main() {
	sayHelloWithFilter("Dani", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Dian", filter)
}