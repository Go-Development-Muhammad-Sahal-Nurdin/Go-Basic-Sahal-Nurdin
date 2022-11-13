// package main

// import "fmt"

// const pi = 22 / 7

// var jari float64

// func main() {
// 	fmt.Printf("Masukkan jari-jari lingkaran = ")
// 	fmt.Scan(&jari)
// 	luas := pi * (jari * jari)
// 	fmt.Println("Luas Lingkaran = ", luas)
// }

package main

import "fmt"

const pi = 22 / 7

var jari int

func main() {
	fmt.Printf("Masukkan jari-jari lingkaran = ")
	fmt.Scan(&jari)
	luas := pi * (jari * jari)
	fmt.Println("Luas Lingkaran = ", luas)
}
