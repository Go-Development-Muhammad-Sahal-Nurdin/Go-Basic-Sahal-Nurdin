package main

import "fmt"

func main() {
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(arr); i++ {
		if arr[i] == 6 {
			fmt.Println("absen", arr[i], "izin")
		} else if arr[i] == 4 {
			(fmt.Println("absen", arr[i], "sakit"))
		} else {
			fmt.Println("absen", arr[i], "hadir")
		}
	}
}
