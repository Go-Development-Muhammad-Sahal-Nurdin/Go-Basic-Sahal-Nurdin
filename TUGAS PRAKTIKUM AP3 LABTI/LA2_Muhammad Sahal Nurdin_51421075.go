// MUHAMMAD SAHAL NURDIN
// 2IA07
// 51421075
package main

import "fmt"

//Ketentuan 1 : Menggunakan Struct
type mahasiswa struct {
	// Deklarasi Variabel tinggi
	tinggi []float64
}

func main() {

	siswa := mahasiswa{}
	// Data tinggi 10 teman saya
	siswa.tinggi = []float64{165, 180, 169, 160, 170, 170, 162, 152, 155, 175}
	inisiasiJumlah := 0

	// Ketentuan 2: Penjumlahan menggunakan perulangan
	for i := 0; i < len(siswa.tinggi); i++ {
		inisiasiJumlah = inisiasiJumlah + int(siswa.tinggi[i])
		// Ini akan mengakses index dari index 0 sampai 10
		if i == 9 {
			// Ketentuan 3: Tambahkan rata-rata menggunakan for
			// Perulangan dibagi rata-rata
			penjumlahanDataTinggi := float64(inisiasiJumlah)
			rataRataTinggi := penjumlahanDataTinggi / 10
			fmt.Printf("Rata-rata tinggi badan mahasiswa: %.1f cm\n", rataRataTinggi)
		}
	}

	fmt.Print("------------------------------------------- \n")
	// Ketentuan 4: akses lokasi index pertama dan keluarkan outputnya
	fmt.Printf("Nilai pada Index pertama adalah: %.0f cm\n", *&siswa.tinggi[0])
	fmt.Printf("Lokasi Storage Index pertama adalah: %d \n", &siswa.tinggi[0])

}
