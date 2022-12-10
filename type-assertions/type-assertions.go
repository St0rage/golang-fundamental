package main

import (
	"fmt"
)

func random() interface{} {
	return true
}

func main() {
	// type assertions
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// type assertions Switch (cara aman)
	switch value := result.(type) {
	case string :
		fmt.Println("Value", value, "is string")
	case int :
		fmt.Println("Value", value, "is int")
	default : 
		fmt.Println("Unknown type")
	}

}