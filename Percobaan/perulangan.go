package main

import "fmt"

func main() {
	var i int = 1
	for i = 1; i <= 7; i++ {
		if i <= 3 {
			fmt.Printf("Absen %d hadir \n", i)
		} else if i == 4 {
			fmt.Printf("Absen %d sakit\n", i)
		} else if i == 5 {
			fmt.Printf("Absen %d hadir \n", i)
		} else if i == 6 {
			fmt.Printf("Absen %d izin \n", i)
		} else if i == 7 {
			fmt.Printf("Absen %d hadir \n", i)
		}
	}
}
