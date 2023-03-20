package main

import "fmt"

func main() {
	var num int
	fmt.Print("Número: ")
	fmt.Scan(&num)
	if num%3 == 0 && num%5 == 0 {
		fmt.Println("O número é simultaneamente divisível por 3 e 5.")
	} else {
		fmt.Println("O número não é simultaneamente divisível por 3 e por 5.")
	}
}
