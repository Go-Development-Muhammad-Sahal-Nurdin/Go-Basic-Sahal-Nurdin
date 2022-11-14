package main

import "fmt"

type student struct {
	name  string
	grade int
}

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, my name is", customer.Name)
}

func main() {
	var s1 student
	s1.name = "Muhammad Sahal Nurdin"
	s1.grade = 90
	fmt.Println("Name	: ", s1.name)
	fmt.Println("Grade	: ", s1.grade)

	var sahal Customer
	sahal.Name = "Muhammad Sahal Nurdin"
	sahal.Address = "Indonesia"
	sahal.Age = 19
	fmt.Println(sahal)
	// Struct literals

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}

	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	rully := Customer{Name: "rully"}
	rully.sayHello()

}
