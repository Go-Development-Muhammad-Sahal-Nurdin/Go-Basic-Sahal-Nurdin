package main

import "fmt"

func main() {
	const firstName1 string = "Nana"
	const lastName1 string = "Komatsu"
	const value = 1000

	fmt.Println(firstName1)
	fmt.Println(lastName1)
	fmt.Println(value)

	//  MULTIPLE CONSTANT
	const (
		firstName string = "Muhammad Sahal Nurdin"
		lastName         = "Nurdin"
		umur             = 18
	)
	// TIDAK BAKAL DICOMPLAIN JIKA TIDAK DI PRINT
	// fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(umur)

}
