package main

import (
	"fmt"
	"strconv"
)

func main() {
	// strconv = string conversion
	// konversi dengan beda tipe data, misal int ke string atau sebaliknya

	// strconv bisa mengembalikan error, sehingga perlu menyiapkan variable error

	// konversi string ke boolean
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	// konversi string ke number
	// Parseint mengembalikan 3 data, (s string, base int, bitSize int)
	number, err := strconv.ParseInt("1000000", 10, 64) // 10 artinya basenya desimal, jika biner 2, oktal 8, dst
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	// konversi dari int ke string
	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	// Itoa = konversi int ke string
	// Atoi = konversi string ke int
	valueStr := strconv.Itoa(3000000)
	fmt.Println(valueStr)
	valueInt, _ := strconv.Atoi("2000000")
	fmt.Println(valueInt)

}
