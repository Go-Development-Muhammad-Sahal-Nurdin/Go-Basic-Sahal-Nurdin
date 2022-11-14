package main

import "fmt"

func main() {
	var a *int
	var b int

	a = &b
	b = 5
	*a = 3

	fmt.Println("alamat a = ", a)
	fmt.Println("nilai a = ", *a)
	fmt.Println("alamat b = ", &b)
	fmt.Println("nilai b = ", b)

}
