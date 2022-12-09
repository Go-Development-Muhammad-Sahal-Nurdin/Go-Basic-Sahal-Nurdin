// // Soal Nomor 1 : Membuat program nested loop

// package main

// import "fmt"

// // func main() {
// // 	fmt.Println("Muhammad Sahal Nurdin")
// // 	fmt.Println("51421075")
// // 	fmt.Println("2IA07")

// // 	// Deklarasi variabel integer
// // 	var banyakBaris int

// // 	// Perintah untuk memasukkan panjang dari iterasi
// // 	fmt.Printf("Masukkan value dari baris yang ada: ")
// // 	// Menangkap input dengan tipe data integer
// // 	fmt.Scanf("%d", &banyakBaris)
// // 	for i := 1; i <= banyakBaris; i++ {
// // 		for j := 1; j <= i; j++ {
// // 			fmt.Print("+ ")
// // 		}
// // 		// New line ke bawah
// // 		fmt.Print("\n")
// // 	}
// // }

// Soal Nomor 2
package main

import (
	"fmt"
	"math"
)

var (
	weight, height, bmi float64
)

func main() {
	fmt.Print("Masukkan berat badan dalam kilogram: ")
	fmt.Scanln(&weight)

	fmt.Print("Masukkan tinggi dalam meter: ")
	fmt.Scanln(&height)

	fmt.Println("Berat Anda : ", weight, "kilogram")
	fmt.Println("Tinggi Anda: ", height, "meter")

	bmi = weight / math.Pow(height, 2)

	fmt.Println("Skor BMI kamu adalah :", bmi)
	fmt.Print("Kategori kamu adalah : ")

	if bmi < float64(18.5) {
		fmt.Println("Kekurangan berat badan")
	} else if bmi < 25 {
		fmt.Println("Berat badan normal")
	} else if bmi < 30 {
		fmt.Println("Kelebihan berat Badan")
	} else {
		fmt.Println("Obesitas")
	}

	//   Menghitung standar bmi yakni 25 poin
	normalWeight := 25 * math.Pow(height, 2)
	delta := weight - normalWeight
	var intNormalWeight int = int(normalWeight)

	fmt.Printf("Berat badan normal untuk tinggi Anda adalah: %0.2v kilogram.\n", intNormalWeight)

	if (delta > 0) && (bmi > 30) {
		fmt.Printf("Kamu perlu mengurang %0.2v kilograms\n", math.Abs(delta))
	}

	if (delta < 0) && (bmi < float64(18.5)) {
		fmt.Printf("Kamu perlu menamabah %0.2v kilograms.\n", math.Abs(delta))
	}

}
