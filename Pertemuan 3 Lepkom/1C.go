// package main

// // Muhammad Sahal Nurdin 51421075

// import "fmt"

// func main() {
// 	for i := 1; i < 11; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i, "adalah Angka Genap")
// 		} else if i%2 == 1 {
// 			fmt.Println(i, "adalah Angka Ganjil")
// 		}

// 	}
// }

package main

// Muhammad Sahal Nurdin 51421075

import "fmt"

func main() {
	var i int = 1
	for i < 11 {
		if i%2 == 0 {
			fmt.Println(i, "adalah Angka Genap")
		} else if i%2 != 0 {
			fmt.Println(i, "adalah Angka Ganjil")
		} else {
			fmt.Println("Sesuatu ada yang salah")
		}
		i = i + 1
	}
}
