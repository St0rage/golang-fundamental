package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Dani")
	data.PushBack("Yudistira")
	data.PushBack("Maulana")
	data.PushFront("Veronica")

	// untuk mengambil data yang selanjutnya
	// data.Front().Next().Next()

	// untuk mengambil data yang sebelumnya
	// data.Back().Prev().Prev()

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	fmt.Print("\n\n")

	// iterasi data list
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Print("\n\n")

	// iterasi reverse
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
