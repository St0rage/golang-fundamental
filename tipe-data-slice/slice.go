package main

import "fmt"

func main() {
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

	var slice1 = months[4:7]

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[0] = "Mei Lagi"
	// fmt.Println(months)

	var slice2 = months[11:]
	fmt.Println(slice2)
	var slice3 = append(slice2, "Dani")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)



	fmt.Println("")
	newSlice := make([]string, 2, 5) // 2 adalah length, 5 adalah capacity
	newSlice[0] = "Dani"
	newSlice[1] = "Yudistira"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// make() = membuat slice baru
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}