package main

import (
	"fmt"
)

func main() {
	var num, sum int
	var avg float64
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
		sum += num
		count++
	}
	avg = float64(sum) / float64(count)
	fmt.Println("A média é:", avg)

}
