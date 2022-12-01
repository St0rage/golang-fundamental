package main

import "fmt"

func main() {
	//deklarasi array
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// cara membuat slice array[low:high]
	// (4) adalah pointer acuan array dan index low [low] {mei}
	// (7) adalah index high (data yang diambil sebelum index high jadi data yang diambil yang ke index 6) [high] {juli}
	// pointer = 4 (mei)
	// length = 3 -> adalah panjang datanya bukan arraynya
	// capacity = 8 (mei - desember)
	var slice1 = months[4:7]

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// slice dan array terkoneksi/reference jika slice datanya dirubah array aslinya akan berubah juga
	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[0] = "Mei Lagi"  //[0] adalah index di slicenya
	// fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)
	// append membuat slice baru dengan menambah data ke posisi terkahir, jika kapasitas penuh maka akan membuat array baru yang artinya slice ini tidak lagi reference ke array awal (month)
	// jika menggunakan append dan memasukan data baru, dan arraynya masih mampu menampung data tersebut maka array tersebut yang dipakai (month)
	// kalo tidak sanggup menampung, maka akan membuat array baru dan tidak reference ke array lama (month)
	// dan jika arraynya masih mampu menampung dan kita rubah datanya, maka semua yang reference/slice yang reference ke array tsb (month) akan kena impactnya
	var slice3 = append(slice2, "Dani")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)



	fmt.Printf("\n\n")
	// make() = membuat slice baru
	// saat menggunakan "make()" sebenarnya ada array didalamnya. (contoh dibawah akan buat array dengan panjang 5)
	newSlice := make([]string, 2, 5) // 2 adalah length (panjang data didalam slice), 5 adalah capacity
	newSlice[0] = "Dani"
	newSlice[1] = "Yudistira"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))



	// copy() copy slice
	// jika slice baru length nya kurang dari slice yang akan dicopy maka data yang diambil adalah sebagian
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1,2,3,4,5} // penulisan deklarasi array
	iniSlice := []int{1,2,3,4,5} //penulisan deklarasi slice

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}