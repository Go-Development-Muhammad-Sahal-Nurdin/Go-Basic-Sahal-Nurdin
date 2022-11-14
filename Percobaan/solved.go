package main

import "fmt"

func main() {
	var indeks int
	for indeks = 1; indeks <= 7; indeks++ {
		switch indeks {
		case 4:
			fmt.Printf("Absen %d sakit\n", indeks)
		case 6:
			fmt.Printf("Absen %d izin\n", indeks)
		default:
			fmt.Printf("Absen %d hadir \n", indeks)
		}
	}
}
