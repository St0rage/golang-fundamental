package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("d([a-z][a-z])i")

	fmt.Println(regex.MatchString("dani"))
	fmt.Println(regex.MatchString("doni"))
	fmt.Println(regex.MatchString("dENIi"))

	search := regex.FindAllString("dani dian doni deni", 10) // mencari yang sama dengan yang MusCompile
	fmt.Println(search)
}
