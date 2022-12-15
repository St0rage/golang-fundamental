package main

import (
	// blank identifier (_), untuk mengatasi golang komplain karena import package tidak digunakan, hal ini juga akan tetap menjalankan "init" nya
	// jika ingin diakses tidak perlu menggunakan (_)
	_ "golang-fundamental/database"
)

func main() {
	// result := database.GetDatabase()
	// fmt.Println(result)
}
