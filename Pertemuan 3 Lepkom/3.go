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
	var i int = 1
	var j int

	for {
		fmt.Print("Input Nilai I: ")
		fmt.Scan(&j)
		i++
		if j%2 == 0 {
			fmt.Println(j, "Adalah Angka Genap")
		} else if j%2 != 0 {
			fmt.Println(j, "Adalah Angka Ganjil")
		} else {
			fmt.Print("Something Wrong!")
		}
		j = j + 1
		if i > 10 {
			fmt.Print("Pertanyaan selesai, Karena angka yang dapat di-input sudah melebihi dari 10. Terima Kasih")
			break

		}
	}

}
