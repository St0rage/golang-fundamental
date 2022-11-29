package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool

	var noKtpDani NoKtp = "123123124128378123"
	var marriedStatus Married = false
	fmt.Println(noKtpDani)
	fmt.Println(marriedStatus)

}