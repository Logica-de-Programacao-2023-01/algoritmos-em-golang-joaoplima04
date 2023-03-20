package main

import "fmt"

func main() {
	var numero, ante, suc int
	fmt.Print("Número: ")
	fmt.Scan(&numero)
	ante = numero - 1
	suc = numero + 1
	fmt.Println("Seu sucessor e antecessor são respectivamente: ", suc, ante)

}
