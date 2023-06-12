//Create a struct house with variables noRooms, price and city
//The golang example creates a new struct. Then it sets the variables.

package main

import "fmt"

type Person struct {
	name string
	job  string
}

type House struct {
	noRooms bool
	price   int
	city    string
}

func main() {
	var aperson Person

	aperson.name = "Albert"
	aperson.job = "Professor"

	fmt.Printf("aperson.name =  %s\n", aperson.name)
	fmt.Printf("aperson.job  =  %s\n\n", aperson.job)

	var inDaHouse House

	inDaHouse.noRooms = true
	inDaHouse.price = 23
	inDaHouse.city = "Accra"

	fmt.Printf("Information on the house\n")
	fmt.Printf("Are there available rooms in the house? %v \n", inDaHouse.noRooms)
	fmt.Printf("The price is Ghc%d\n", inDaHouse.price)
	fmt.Printf("The house is located in %v \n", inDaHouse.city)
}
