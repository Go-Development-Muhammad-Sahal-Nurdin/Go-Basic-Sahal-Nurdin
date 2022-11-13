package main

// Muhammad Sahal Nurdin 51421075

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}

}
