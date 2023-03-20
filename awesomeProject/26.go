package main

import "fmt"

func main() {
	var num, tabuada int
	fmt.Print("Digite um número de 1 a 9: ")
	fmt.Scan(&num)
	tabuada = 0
	for tabuada != 10 {
		tabuada++
		fmt.Println(num * tabuada)
	}
}
