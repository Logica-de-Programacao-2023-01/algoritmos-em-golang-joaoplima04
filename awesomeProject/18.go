package main

import "fmt"

func main() {
	var num int
	fmt.Print("digite um número entre 1 a 7: ")
	fmt.Scan(&num)
	if num == 1 {
		fmt.Println("Segunda")
	}
	if num == 2 {
		fmt.Println("Terça")
	}
	if num == 3 {
		fmt.Println("Quarta")
	}
	if num == 4 {
		fmt.Println("Quinta")
	}
	if num == 5 {
		fmt.Println("Sexta")
	}
	if num == 6 {
		fmt.Println("Sábado")
	}
	if num == 7 {
		fmt.Println("Domingo")
	}

}
