package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke-", counter)
	// 	counter++
	// }

	// FOR DENGAN STATEMENT
	// INIT STATEMENT POST STATEMENT
	for counter := 1;counter <= 10;counter++ {
		fmt.Println("Perulangan ke-", counter)
	}

	// AKSES SLICE
	slice := []string{"Muhammad", "Sahal", "Nurdin"}
	for i:=0; i< len(slice); i++{
		fmt.Println(slice[i])
	}

	// MENGGUNAKAN FOR RANGE UNTUK SLICE
	// names := []string{"Nana", "Komatsu"}
	// for i, value:= range names {
	// 	fmt.Println("Index", i, "=", value)
	// 	fmt.Println()
	// }

	// PRINT ISINYA SAJA BUKAN INDEKSNYA PAKAI UNDERSCORE
	names := []string{"Nana", "Komatsu"}
	for _, value:= range names {
		fmt.Println(value)
	}

	// FOR UNTUK MAP
	person := make(map[string]string) 
	person["name"] = "Nana Komatsu"
	person["title"] = "artist"
	for key, value:= range person {
		fmt.Println(key, "=", value)
	}
}