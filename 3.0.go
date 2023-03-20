package main

import "fmt"

func main() {
	var base float64
	var altura float64
	var profundidade float64
	fmt.Print("Qual é a base? ")
	fmt.Scan(&base)
	fmt.Println("Qual é a altura? ")
	fmt.Scan(&altura)
	fmt.Println("Qual a profundidade? ")
	fmt.Scan(&profundidade)
	fmt.Println("O volume da caixa é :", base*altura*profundidade)

}
