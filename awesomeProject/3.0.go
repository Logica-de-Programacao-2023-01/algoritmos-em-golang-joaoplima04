package main

import "fmt"

func main() {
	var peso, altura float64
	fmt.Print("Qual o seu peso? ")
	fmt.Scan(&peso)
	fmt.Println("Qual a sua altura? ")
	fmt.Scan(&altura)
	fmt.Println("seu IMC é:", peso/(altura*altura))

}
