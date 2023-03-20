package main

import "fmt"

func main() {
	var peso_kg float64
	var peso_lbs float64
	fmt.Print("Qual o seu peso em kg? ")
	fmt.Scan(peso_kg)
	peso_lbs = peso_kg / 0.45
	fmt.Println("Seu peso em lbs é:", peso_lbs)

}