package main

import "fmt"

const pi = 3.14
var jari int


func main() {
	fmt.Printf("Masukkan jari-jari lingkaran = ")
	fmt.Scan(&jari)
	luas := pi * (jari * jari)
	fmt.Println("Luas Lingkaran = "luas)
}
