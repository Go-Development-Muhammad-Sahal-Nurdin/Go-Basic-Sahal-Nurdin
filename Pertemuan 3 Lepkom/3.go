// package main

//  Muhammad Sahal Nurdin 51421075

// import "fmt"

// func main() {
// 	var i, inputNilai int

// 	for i = 0; i < 10; i++ {

// 		fmt.Print("Input Nilai I: ")
// 		fmt.Scan(&inputNilai)
// 		if inputNilai%2 == 0 {
// 			fmt.Println(inputNilai, "Adalah Angka Genap")
// 		} else if inputNilai%2 != 0 {
// 			fmt.Println(inputNilai, "Adalah Angka Ganjil")
// 		}

// 	}

// }

package main

// Muhammad Sahal Nurdin 51421075

import "fmt"

func main() {
	var i, inputNilai int

	for i = 0; i < 10; i++ {

		fmt.Print("Input Nilai I: ")
		fmt.Scan(&inputNilai)
		if inputNilai%2 == 0 {
			fmt.Println(inputNilai, "Adalah Angka Genap")
		} else if inputNilai%2 != 0 {
			fmt.Println(inputNilai, "Adalah Angka Ganjil")
		}

	}

}
