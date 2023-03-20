package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int

	fmt.Println("Primeioro número: ")
	fmt.Scan(&num1)
	fmt.Println("Segundo: ")
	fmt.Scan(&num2)
	fmt.Println("Terceiro: ")
	fmt.Scan(&num3)

	maior := num1
	if num2 > maior {
		maior = num2
	}
	if num3 > maior {
		maior = num3
	}

	fmt.Println("O maior número é:", maior)
}
