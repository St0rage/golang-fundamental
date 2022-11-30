package main

import "fmt"

func main() {
	var name = "Dani Yudistira"

	if name == "Dani" {
		fmt.Println("Hello Dani")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hi, Kenalan donk")
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}