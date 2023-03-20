package main

import "fmt"

func main() {
	var num_1, num_2, num_3 float64
	var media float64
	fmt.Print("digite três números: ")
	fmt.Scan(&num_1, &num_2, &num_3)
	media = ((num_1 * 2) + (num_2 * 3) + (num_3 * 5)) / 10
	fmt.Println("A média com os pesos 2,3 e 5 é:", media)

}
