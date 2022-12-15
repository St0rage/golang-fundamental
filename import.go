package main

import (
	"fmt"
	"golang-fundamental/helper"
)

func main() {
	helper.SayHello("Dani")
	// helper.sayGoodbye("Dani") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}
