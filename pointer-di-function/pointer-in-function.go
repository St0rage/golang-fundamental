package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var alamat = Address{
		City: "Garut",
		Province: "Jawa Barat",
		Country: "",
	}
	var alamatPointer *Address = &alamat
	// Contoh 1
	// ChangeCountryToIndonesia(&alamat)
	// Contoh 2
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)

	/** 
		disarankan jika mempunyai struct dengan field yang banyak disarankan pada function menggunakan pointer, 
		karena jika tidak akan selalu diduplikasi sehingga penggunaan memori menjadi bengkak
	*/
}