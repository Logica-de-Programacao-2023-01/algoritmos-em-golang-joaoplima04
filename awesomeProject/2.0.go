package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Número: ")
	fmt.Scan(&numero)
	fmt.Println("O dobro é:", numero*2)
	fmt.Println("o triplo é:", numero*3)
	fmt.Println("O quadruplo é:", numero*4)

}
