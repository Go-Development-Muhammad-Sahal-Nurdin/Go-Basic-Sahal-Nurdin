package main

import "fmt"

func main() {

	var data = map[string]mahasiswa{
		"51421075": {
			"MUHAMMAD SAHAL NURDIN",
			"2IA07",
		},
		"10115961": {
			"BUDI ANTONO",
			"4KA20",
		},
	}
	var search string
	fmt.Print("Masukkan NPM anda ? ")
	fmt.Scanf("%s", &search)

	var value, ok = data[search]

	if ok {
		fmt.Printf("Nama anda %s \nkelas anda %s", value.Nama, value.Kelas)
	} else {
		fmt.Println("data tidak ditemukan")
	}
}

type mahasiswa struct {
	Nama  string
	Kelas string
}
