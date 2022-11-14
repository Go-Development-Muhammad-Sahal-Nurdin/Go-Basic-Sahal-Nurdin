package main

import "fmt"

func main() {
	// If else
	var point = 9
	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Println("tidak lulus, nilai anda %d\n", point)
	}

	// switch
	_ = "hallo"
	var point1 = 6
	switch point1 {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not used")
	}

	// LA
	var i int
	for i = 1; i <= 5; i++ {
		fmt.Println("Hello World")
	}

	// Program absensi
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
