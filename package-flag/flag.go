package main

import (
	"flag"
	"fmt"
)

func main() {
	/**
	Parameter ke-1 nama flagnya
	Parameter ke-2 default valuenya
	Parameter ke-3 catatan/cara pengunaan

	bentunya dari flag.String ini merupakan pointer
	*/
	var host *string = flag.String("host", "localhost", "Put your host")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "root", "Put your database user")
	var number *int = flag.Int("number", 100, "Put your number")

	// melakukan parsing
	flag.Parse()

	// karena flag hasilnya pointer maka harus menggunakan (*)
	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Password : ", *password)
	fmt.Println("Number : ", *number)
}
