package main

import "fmt"

func sum(arr []int) int {
	result := 0
	for _, i := range arr {
		result += i
	}
	return result
}
func main() {

	tinggiBadan := []int{169, 165, 169, 160, 170}
	first := tinggiBadan[0]
	var p *int
	p = &first
	fmt.Println("Given list is: ", tinggiBadan)
	res := sum(tinggiBadan)
	fmt.Println("Average of numbers is: ", res/len(tinggiBadan))
	fmt.Println(first)
	fmt.Println(p)
}
