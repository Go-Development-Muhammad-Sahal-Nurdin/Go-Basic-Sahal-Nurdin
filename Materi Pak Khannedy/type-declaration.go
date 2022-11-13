package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpSahal NoKTP = "73468263478"
	fmt.Println(noKtpSahal)
	var marriedStatus Married = false
	fmt.Println(marriedStatus)
}
