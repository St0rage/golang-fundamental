package main

import "fmt"

func getGoodBy(name string) string {
	return "Good Bye " + name
}

func main() {

	sayGoodBy := getGoodBy

	result := sayGoodBy("Dani")
	fmt.Println(result)

}