package main

import "fmt"

func main() {
	var name = "Sahal"
	if name == "Sahal" {
	fmt.Println("Hello, Sahal")
	} else if name == "Nurdin" {
		fmt.Println("Hello, Nurdin")
	}else {
		fmt.Println("Hai, kenalan donk")
	}
	// IF SHORT STATEMENT CUMA ADA DI GOLANG
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}