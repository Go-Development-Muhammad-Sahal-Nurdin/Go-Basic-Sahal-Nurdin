package main

import "fmt"

func main() {
	var ekspektasiGaji int
	var gajiSekarang *int = &ekspektasiGaji

	fmt.Print("Masukkan gaji anda: ")
	fmt.Scan(gajiSekarang)

	fmt.Print("Masukkan gaji anda inginkan : ")
	fmt.Scan(&ekspektasiGaji)

	naikanGaji(gajiSekarang, ekspektasiGaji)

	fmt.Printf("\nGaji anda sekarang %d\n", *gajiSekarang)
}

func naikanGaji(gajiAwal *int, gajiaAkhir int) {
	*gajiAwal = gajiaAkhir
}
