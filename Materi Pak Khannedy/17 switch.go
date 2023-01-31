package main

import "fmt"

func main() {
	// SWITCH DIGUNAKAN UNTUK OPERASI KONDISI YANG SEDERHANA SAJA
	name := "Sahal"
	switch name {
	case "Sahal":
		fmt.Println(("Hallo Sahal"))
		fmt.Println(("Hallo Sahal"))
	case "Nurdin":
		fmt.Println(("Hallo Nurdin"))
		fmt.Println(("Hallo Nurdin"))
	default:
		fmt.Println(("Kenalan donk"))
		fmt.Println(("Kenalan donk"))
	}

	//  SHORT STATEMENT
	switch length := len(name); length > 5  {
	case true:
		fmt.Println(("Nama terlalu panjang"))
	case false:
		fmt.Println(("Nama sudah benar"))
	}

	// SWITCH TANPA KONDISI
	length := len(name)
	switch  {
	case length > 10 :
		fmt.Println(("Nama terlalu panjang"))
	case length > 5:
		fmt.Println(("Nama lumayan panjang"))
	default:
		fmt.Println("Nama sudah benar")
	}
}