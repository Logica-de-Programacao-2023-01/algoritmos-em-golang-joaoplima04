package main

import "fmt"

func main() {
	var preco, precod float64
	fmt.Print("Qual o preço do produto? ")
	fmt.Scan(&preco)
	precod = preco*0.1 + preco
	fmt.Println("O preço com 10% de desconto é:", precod)

}
