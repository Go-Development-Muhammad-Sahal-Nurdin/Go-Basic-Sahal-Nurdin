// Soal Nomor 1 : Membuat program nested loop

package main

import "fmt"

func main() {
	fmt.Println("Muhammad Sahal Nurdin")
	fmt.Println("51421075")
	fmt.Println("2IA07")

	// Deklarasi variabel integer
	var banyakBaris int

	// Perintah untuk memasukkan panjang dari iterasi
	fmt.Printf("Masukkan value dari baris yang ada: ")
	// Menangkap input dengan tipe data integer
	fmt.Scanf("%d", &banyakBaris)
	for i := 1; i <= banyakBaris; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("+ ")
		}
		// New line ke bawah
		fmt.Print("\n")
	}
}
