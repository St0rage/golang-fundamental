package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// pass by value = data didalamnya diduplikat dan dikirim ke tempat lain/lokasi memori yang lain
	// pass by reference = data nya tetap ditempat/lokasi yang sama/lokasi memori yang sama
	// address1 := Address{"Garut", "Jawa Barat", "Indonesia"}
	// address2 := address1

	// ketika address2.City diubah adress1.City tidak akan berubah karena pass by value/datanya di duplikat sehingga tidak reference ke address1
	// address2.City = "Bandung"

	// fmt.Println(address1)
	// fmt.Println(address2)

	// untuk membuat variable dengan nilai pointer ke variable yang lain menggunakan operator(&) diikuti dengan nama variablenya
	address1 := Address{"Garut", "Jawa Barat", "Indonesia"}
	address2 := &address1
	// kode diatas sama dengan dibawah, bintang (*) artinya pointer
	// var address1 Address = Address{"Garut", "Jawa Barat", "Indoneisa"}
	// var address2 *Address = &address1 


	// pembuktian pindah pointer
	address3 := &address1

	// ketika address2.City diubah address1.City juga akan berubah karena kita menggunakan pointer (&) ke variable address1 sehingga bentuknya pass by reference/mengguunakan data yang sama
	address2.City = "Bandung"

	// kode dibawah address2 adalah pointer dengan menambahkan/assing baru address2 dengan &Address{} -
	// maka akan membuat data di memori baru dan address2 pointer nya mengarah ke struct/memori baru tsb (pindah pointer), sedangkan address1 tetap di struct/memori yang sama sehingga address1 datanya tidak ikut berubah
	// address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}

	// sedangkan kode dibawah adalah memindahkan variabel yang pointernya semula di memori awal ke memori yang baru, sehingga address1 pindah pointernya reference ke address2
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	
	fmt.Print("\n\n")

	// New
	// kode dibawah langsung pointer ke address/struct/memori baru
	var address4 *Address = new(Address) // akan membuat pointer/mengarahkan pointernya ke lokasi memori baruyang isinya data kosong
	address4.City = "Jakarta"
	fmt.Println(address4)

}