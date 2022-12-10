package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return 2
	} else {
		return "Ups"
	}
}

func main() {
	// return value harus interface kosong
	var data interface{} = Ups(1)
	fmt.Println(data)
}

// interface kosong (interface{}) tidak peduli dengan tipe data apapun