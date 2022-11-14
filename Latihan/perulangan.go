package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	fmt.Println("Hanya Kondisi")
	i := 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	fmt.Println("Tanpa Argumen")
	var x = 1
	for {
		fmt.Println("Angka", x)
		x++
		if x == 5 {
			break
		}
	}

	fmt.Println("Basic")
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	fmt.Println("Kode Program For Range")
	names := []string{"Muhammad", "Sahal", "Nurdin"}
	for index, name := range names {
		//fmt.Println("index", index, "=", name)
		fmt.Println(index, name)
	}

	for i := 0; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}

	for i := 0; i <= 5; i++ {
		for j := 1; j <= 5-i; j++ {
			fmt.Print("+ ")
		}
		fmt.Print("\n")
	}

}
