package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhammad"
	names[1] = "Sahal"
	names[2] = "Nurdin"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Membuat array langsung
	// jika array sudah dibuat ukurannya maka tidak bisa ditambahkan lagi elemennya

	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string
	// len bukan untuk banyak data
	fmt.Println(len(lagi))

}
