package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("digite um número: ")
	fmt.Scanln(&num1)
	fmt.Println("Digite outro número: ")
	fmt.Scan(&num2)
	if num1 > 0 && num2 > 0 {
		fmt.Println("A multiplicação entre eles é:", num2*num1)
	} else {
		fmt.Println("A soma entre eles é:", num2+num1)
	}
}
