package main

import "fmt"

func main() {
	var altura float64
	var base float64
	fmt.Print("Base: ")
	fmt.Scan(&base)
	fmt.Println("Altura: ")
	fmt.Scan(&altura)
	fmt.Println("Área = ", base*altura)
}
