package main

import "fmt"

func main() {
	//var
	var firstName string = "John"
	var firstName_1 = "John"
	firstName_2 := "John"

	fmt.Println(firstName)
	fmt.Println(firstName_1)
	fmt.Println(firstName_2)

	//multiple variable
	var lastName, address string = "Doe", "USA"
	var lastName_1, address_1 = "Doe", "USA"
	lastName_2, address_2 := "Doe", "USA"
	var (
		lastName_3 = "Doe"
		address_3  = "USA"
	)

	fmt.Println(lastName, address)
	fmt.Println(lastName_1, address_1)
	fmt.Println(lastName_2, address_2)
	fmt.Println(lastName_3, address_3)

	//constant
	const birthDate string = "2000-01-01"
	const birthPlace = "Jakarta"
	const (
		statusOK = 200
		notFound = 404
	)

	fmt.Println(birthDate)
	fmt.Println(birthPlace)
	fmt.Println(statusOK)
}
