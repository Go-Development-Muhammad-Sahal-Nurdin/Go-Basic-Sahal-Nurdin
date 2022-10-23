package main

import "fmt"

func main() {
	var i int = 1
	for i = 1; i <= 7; i++ {
		switch i {
		case 4:
			fmt.Printf("Absen %d sakit\n", i)
		case 6:
			fmt.Printf("Absen %d izin\n", i)
		default:
			fmt.Printf("Absen %d hadir \n", i)
		}
	}
}
