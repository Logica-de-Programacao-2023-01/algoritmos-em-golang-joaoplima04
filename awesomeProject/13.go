package main

import "fmt"

func main() {
	var num int
	fmt.Print("Numero: ")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("O número é par.")
	} else {
		println("O número é ímpar.")
	}
}
