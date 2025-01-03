package main

import "fmt"

func main() {
	//integer
	//tipe data angka termasuk minus
	var positiveNumber int = 89
	var negativeNumber int = -124342

	fmt.Println()
	fmt.Println("type data integer")
	fmt.Println(positiveNumber)
	fmt.Println(negativeNumber)

	//unsigned integer
	//tipe data angka tidak termasuk minus dimulai dari 0
	var unsignedNumber uint = 89

	fmt.Println()
	fmt.Println("type data unsigned integer")
	fmt.Println(unsignedNumber)

	//float
	//tipe data angka desimal
	var decimalNumber float32 = 2.62

	fmt.Println()
	fmt.Println("type data float")
	fmt.Println(decimalNumber)

	//alias
	//byte = uint8
	//rune = int32
	//int = minimal 32 bit
	//uint = minimal 32 bit
}
