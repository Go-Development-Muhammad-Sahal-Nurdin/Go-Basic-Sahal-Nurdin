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
	fmt.Println("Berikut ini adalah list dari tinggi badan teman saya: ", tinggiBadan)
	res := sum(tinggiBadan)
	fmt.Println("Rata-rata tinggi badan adalah: ", res/len(tinggiBadan))
	fmt.Println(first)
	fmt.Println(p)
}
