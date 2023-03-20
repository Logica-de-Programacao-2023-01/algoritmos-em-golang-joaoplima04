package main

import "fmt"

func main() {
	var diast, diaria, salario int
	fmt.Print("Quantos dias o funcionário trabalhou? ")
	fmt.Scan(&diast)
	fmt.Println("Qual é o valor da sua diária? ")
	fmt.Scan(&diaria)
	salario = diast * diaria
	fmt.Println("Seu salário é:", salario)

}
