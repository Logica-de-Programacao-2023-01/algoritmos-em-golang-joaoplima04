package main

import "fmt"

func main() {
	var idade_anos, idade_dias int
	fmt.Print("Qual a sua idade? ")
	fmt.Scan(&idade_anos)
	idade_dias = idade_anos * 365
	fmt.Println("A sua idade em dias é:", idade_dias)

}
