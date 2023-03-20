package main

import "fmt"

func main() {
	var salario float64
	var aumento float64
	fmt.Print("Salário: ")
	fmt.Scan(&salario)
	aumento = salario*0.15 + salario
	fmt.Println("Seu novo salário com o aumento de 15% é:", aumento)

}
