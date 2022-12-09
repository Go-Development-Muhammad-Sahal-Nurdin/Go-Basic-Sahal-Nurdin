package main

import "fmt"

var namaUmurMap map[string]int

func main() {
	// Inisialisasi Mapping
	namaUmurMap = map[string]int{
		"Muhammad Sahal Nurdin": 19,
		"Nana":                  26,
	}
	fmt.Println("Mencetak Umur Sahal:", namaUmurMap["Muhammad Sahal Nurdin"])
	// Pembacaan Map
	for key, value := range namaUmurMap {
		fmt.Printf("%v berumur %d tahun\n", key, value)
	}
}
