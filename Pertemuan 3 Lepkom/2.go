// package main

// // Muhammad Sahal Nurdin 51421075

// import "fmt"

// var nilaiUTS, nilaiUAS, nilaiTotal, UTS, UAS float64

// func main() {
// 	// Input user
// 	fmt.Print("Masukkan nilai UTS: ")
// 	fmt.Scan(&UTS)
// 	fmt.Print("Masukkan nilai UAS: ")
// 	fmt.Scan(&UAS)

// 	// Perhitungan bobot
// 	nilaiUTS = 0.7 * UTS
// 	nilaiUAS = 0.3 * UAS
// 	nilaiTotal = nilaiUTS + nilaiUAS
// 	fmt.Println("Total Nilai = ", nilaiTotal)

// 	// Logic
// 	if nilaiTotal <= 20 {
// 		fmt.Print("Grade e")
// 	} else if nilaiTotal >= 21 && nilaiTotal <= 40 {
// 		fmt.Print("Grade D")
// 	} else if nilaiTotal >= 41 && nilaiTotal <= 60 {
// 		fmt.Print("Grade C")
// 	} else if nilaiTotal >= 61 && nilaiTotal <= 80 {
// 		fmt.Print("Grade B")
// 	} else if nilaiTotal >= 81 && nilaiTotal <= 100 {
// 		fmt.Print("Grade A")
// 	} else {
// 		fmt.Print("Nilai tidak ditemukan")
// 	}
// }

package main

// Muhammad Sahal Nurdin 51421075

import "fmt"

func main() {
	var uts, uas, total_nilai float64
	fmt.Print("Masukkan Nilai UTS : ")
	fmt.Scan(&uts)
	fmt.Print("Masukkan Nilai UAS : ")
	fmt.Scan(&uas)

	fmt.Print("\n Nilai UTS = ", uts)
	fmt.Print("\n Nilai UAS = ", uas)

	total_nilai = (0.7 * uts) + (0.3 * uas)
	fmt.Println("\n Total Nilai = ", total_nilai)

	// Logic
	if total_nilai <= 20 {
		fmt.Print("\n Grade E")
	} else if total_nilai <= 40 {
		fmt.Print("\n Grade D")
	} else if total_nilai <= 60 {
		fmt.Print("\n Grade C")
	} else if total_nilai <= 80 {
		fmt.Print("\n Grade B")
	} else {
		fmt.Print("\n Grade A")
	}
}
