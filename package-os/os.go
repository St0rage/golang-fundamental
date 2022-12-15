package main

import (
	"fmt"
	"os"
)

func main() {
	// mengambil argumen yang di input pada terminal, berupa []string/slice
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)

	// cara mengakses argumen, diawali indeks ke 1
	// indeks ke 0 digunakan sistem/lokasi penyimpanan build
	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])

	// ====
	fmt.Print("\n\n")

	// mengambil data hostname, terdapat dua return
	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	// ====
	fmt.Print("\n\n")

	// mengambil data env
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Println(username)
	fmt.Println(password)

}
