package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	// AUGMENTED ASSIGNMENTS
	var i = 10
	i += 10 // I = I + 10
	fmt.Println(i)

	// UNARY OPERATOR
	i++
	fmt.Println(i)
	var negative = -100
	var positive = +100
	fmt.Println(negative)
	fmt.Println(positive)

}
