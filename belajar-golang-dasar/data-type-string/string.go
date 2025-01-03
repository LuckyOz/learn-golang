package main

import "fmt"

func main() {
	// string
	fmt.Println("Belajar Golang Dasar")
	fmt.Println("Belajar" + " " + "Golang" + " " + "Dasar")

	//Function Bawaan String
	// len() => menghitung jumlah karakter pada string
	fmt.Println(len("Belajar Golang Dasar")) // 20

	//string[number] => mengambil karakter pada posisi tertentu
	fmt.Println("Belajar Golang Dasar"[0]) // 66 => B

}
