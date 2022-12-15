package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Dani Yudistira", "Dani"))
	fmt.Println(strings.Contains("Dani Yudistira", "Maulana"))

	fmt.Println(strings.Split("Dani Yudistira", " "))

	fmt.Println(strings.ToLower("Dani Yudistira Maulana"))
	fmt.Println(strings.ToUpper("Dani Yudistira Maulana"))
	fmt.Println(strings.ToTitle("dani yudistira Maulana"))

	fmt.Println(strings.Trim("     Dani Yudistira     ", " "))

	fmt.Println(strings.ReplaceAll("Dani Dani Dani Dani", "Dani", "Dian"))
}
