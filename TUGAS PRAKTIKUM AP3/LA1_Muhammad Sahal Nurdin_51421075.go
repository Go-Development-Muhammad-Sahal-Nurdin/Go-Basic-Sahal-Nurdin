// Soal Nomor 1

package main

import "fmt"

func main() {
	var n int
	fmt.Printf("Masukkan Value: ")
	fmt.Scanf("%d", &n)
	for i := 1; 1 <= n; i++ {
		for j := 1; j <= i; j++ {
		fmt.Print("+ ")
		}
		fmt.Print("\n")
	}
}