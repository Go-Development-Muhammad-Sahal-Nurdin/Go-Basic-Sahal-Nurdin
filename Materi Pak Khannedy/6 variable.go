package main

import "fmt"

func main() {
	var name string

	name = "Muhammad Sahal Nurdin"
	fmt.Println(name)

	name = "Muhammad Sahal Nurdin"
	fmt.Println(name)

	var friendName = "Budi"
	fmt.Println(friendName)

	var age int8 = 19
	fmt.Println(age)

	// TYPE INFERENCE (VAR TIDAK WAJIB)
	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "Muhammad"
		middleName = "Sahal"
		lastName  = "Nurdin"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

}
