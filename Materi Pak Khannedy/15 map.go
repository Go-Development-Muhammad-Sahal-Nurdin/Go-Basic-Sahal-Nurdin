package main
import "fmt"

func main() {
	person := map[string]string{
		"name" : "Sahal",
		"address" : "Jonggol",
	}
	// TAMBAH KEY AND VALUE
	person["title"] = "Programmer"
 
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// BUAT MAP BARU
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Sahal"
	book["ups"] ="Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))

}
