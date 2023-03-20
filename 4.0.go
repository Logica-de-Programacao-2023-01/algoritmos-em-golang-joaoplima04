package main

import "fmt"

func main() {
	var numero_1 int
	var numero_2 int
	var numero_3 int
	var numero_4 int
	fmt.Print("Qual é o primeiro número? ")
	fmt.Scan(&numero_1)
	fmt.Print("Qual é o segundo número? ")
	fmt.Scan(&numero_2)
	fmt.Print("Qual é o terceiro número? ")
	fmt.Scan(numero_3)
	fmt.Print("Qual é o quarto número? ")
	fmt.Scan(&numero_4)
	fmt.Print("A média é: ", (numero_1+numero_2+numero_3+numero_4)/4)
}
