package main

import "fmt"

func main() {
	var salario, aumento float64
	fmt.Print("Qual o seu salário? ")
	fmt.Scan(&salario)
	if salario < 1000 {
		aumento = salario*0.1 + salario
	}
	if salario >= 1000 {
		aumento = salario*0.05 + salario
	}
	fmt.Println("Seu salário com o aumento é:", aumento)
}
