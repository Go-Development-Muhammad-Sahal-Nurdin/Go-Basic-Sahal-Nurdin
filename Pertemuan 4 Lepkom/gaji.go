package main

import "fmt"

func main() {
	var gajiSekarang, ekspektasiGaji int

	fmt.Print("Masukkan gaji anda: ")
	fmt.Scan(&gajiSekarang)

	fmt.Print("Masukkan gaji anda inginkan : ")
	fmt.Scan(&ekspektasiGaji)

	naikanGaji(gajiSekarang, ekspektasiGaji)

	fmt.Printf("\nGaji anda sekarang %d\n", gajiSekarang)
}

func naikanGaji(gajiAwal int, gajiaAkhir int) {
	gajiAwal = gajiaAkhir
}
