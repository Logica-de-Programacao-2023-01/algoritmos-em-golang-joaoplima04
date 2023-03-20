package main

import (
	"fmt"
)

func main() {
	var num, soma int
	var media float64
	count := 0
	fmt.Println("Digite uma lista de números inteiros os separando com espaço, quando terminar digite 0")
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("Erro, Tente novamente")
			continue
		}
		if num == 0 {
			break
		}
		soma += num
		count++
	}
	media = float64(soma) / float64(count)
	fmt.Println("A média é:", media)

}
