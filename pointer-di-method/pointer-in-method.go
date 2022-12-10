package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	// println("Di Method", man.Name)
}

func main() {
	dani := Man{"Dani"}
	dani.Married()
	// pada struct untuk menggunakan pointer tidak perlu menggunakan (&)
	fmt.Println(dani.Name)
}