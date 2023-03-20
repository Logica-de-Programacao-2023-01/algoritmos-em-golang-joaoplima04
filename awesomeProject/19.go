package main

import "fmt"

func main() {
	var num1, num2, num3 float64
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)
	fmt.Println("Segundo: ")
	fmt.Scan(&num2)
	fmt.Println("Terceiro: ")
	fmt.Scan(&num3)
	if num1 > num2 && num2 > num3 {
		fmt.Println("Os números em ordem crescente são: ", num3, num2, num1)
	}
	if num1 > num3 && num2 < num3 {
		fmt.Println("Os números em ordem crescente são: ", num2, num3, num1)
	}
	if num2 > num1 && num1 > num3 {
		fmt.Println("Os números em ordem crescente são: ", num3, num1, num2)
	}
	if num2 > num3 && num3 > num2 {
		fmt.Println("Os números em ordem crescente são: ", num2, num3, num1)
	}
	if num3 > num2 && num2 > num1 {
		fmt.Println("Os números em ordem crescente são: ", num1, num2, num3)
	}
	if num3 > num1 && num2 < num1 {
		fmt.Println("Os números em ordem crescente são: ", num2, num1, num3)
	}

}
