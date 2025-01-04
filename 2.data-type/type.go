package main

import "fmt"

func main() {
	// Tipe data number
	fmt.Println("Satu", 1)
	fmt.Println("Dua koma lima", 2.5)

	// Tipe data boolean
	fmt.Println(true)
	fmt.Println(false)

	// Tipe data string
	fmt.Println("Hello, World!")
	fmt.Println("Jumlah karakter dari kata 'Hello, World!': ", len("Hello, World!"))
	fmt.Println("Byte karakter 'H' dari kata 'Hello, World!': ", "Hello, World!"[0])

	// Cek tipe data
	fmt.Println("Tipe data dari 1: ", fmt.Sprintf("%T", 1))
}
