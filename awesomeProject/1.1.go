package main

import "fmt"

func main() {
	var numero_1 int
	var numero_2 int
	var numero_3 int
	fmt.Print("número 1: ")
	fmt.Scan(&numero_1)
	fmt.Print("numero 2: ")
	fmt.Scan(&numero_2)
	fmt.Print("numero 3: ")
	fmt.Scan(&numero_3)
	fmt.Println("A soma dos 3 números é:", numero_1+numero_2+numero_3)
}
