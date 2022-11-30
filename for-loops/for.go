package main

import (
	"fmt"
)

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan Ke", counter)
	// 	counter++
	// }

	// for dengan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	// for dan array/slice/map
	slice := []string{"Dani", "Yudistira", "Maulana"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for range (khusus array/slice/map)
	// gunakan "_" (underscore) untuk memberitahu golang variabel tidak dibutuhkan
	for _, value := range slice {
		// fmt.Println("Index", i, "=", value)
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Dani"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}