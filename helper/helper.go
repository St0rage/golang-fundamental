package helper

import "fmt"

// tidak bisa diakses dari luar karena tidak diwali huruf besar
var version = 1

// bisa diakses dari luar karena diawali dengan huruf besar
var Application = "Belajar Golang"

// bisa diakses dari luar karena diawali dengan huruf besar pada nama functionnya
func SayHello(name string) {
	fmt.Println("Hello", name)
}

// tidak bisa diakses dari luar karena tidak diawali dengan huruf besar
func sayGoodbye(nama string) {
	fmt.Println("Goodbye", nama)
}
