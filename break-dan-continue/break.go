package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		// jika di break perulangan benar2 berhenti
		fmt.Println("Perulangan ke", i)
	}
}