package main

import "fmt"

func main() {
	var altura, peso, imc float64
	var sexo string
	fmt.Print("Qual o seu sexo? ")
	fmt.Scan(&sexo)
	fmt.Println("Qual a sua altura? ")
	fmt.Scan(&altura)
	fmt.Println("Qual o seu peso? ")
	fmt.Scan(&peso)
	imc = peso / (altura * altura)
	if imc >= 20 && imc <= 24 {
		fmt.Println("Voçê está dentro do peso ideal.")
	}
	if imc >= 25 && imc <= 29 {
		fmt.Println("Voçê está em excesso de peso.")
	}
	if imc >= 30 && imc <= 35 {
		fmt.Println("Voçê está obeso")
	}
	if imc > 35 {
		fmt.Println("Voçê está super obeso.")
	}

}
