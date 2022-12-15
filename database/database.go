package database

import "fmt"

var connection string

// function yang akan otomatis dijalankan ketika package 'database' dipanggil
func init() {
	fmt.Println("Init Dipanggil")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
