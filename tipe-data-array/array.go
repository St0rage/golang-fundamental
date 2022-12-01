package main

import "fmt"

func main() {
	// membuat array
	var names [3]string

	names[0] = "Dani"
	names[1] = "Yudistira"
	names[2] = "Maulana"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// array langsung diset / langsung deklarasi
	var values = [3]int{90,95,80}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// mengambil panjang array bukan panjang isi/data arraynya
	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string

	fmt.Println(len(lagi)) //hasil nya 10
}