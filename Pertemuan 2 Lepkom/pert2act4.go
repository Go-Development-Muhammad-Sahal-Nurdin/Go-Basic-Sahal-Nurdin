// package main

// import "fmt"

// var nilai1, nilai2 float64

// func main() {
// 	defer fmt.Println("--Selesai--")
// 	fmt.Print("Masukkan bilangan 1: ")
// 	fmt.Scan(&nilai1)
// 	fmt.Print("Masukkan bilangan 2: ")
// 	fmt.Scan(&nilai2)
// 	hasil := nilai1 / nilai2
// 	fmt.Printf("Hasil dari Nilai2 / Nilai2 = %.3f\n", hasil)
// }

package main

import "fmt"

var nilai1, nilai2 int

func main() {
	defer fmt.Println("--Selesai--")
	fmt.Print("Masukkan bilangan 1: ")
	fmt.Scan(&nilai1)
	fmt.Print("Masukkan bilangan 2: ")
	fmt.Scan(&nilai2)
	hasil := nilai1 % nilai2
	fmt.Printf("Hasil dari Nilai2 / Nilai2 = %d\n", hasil)
}
