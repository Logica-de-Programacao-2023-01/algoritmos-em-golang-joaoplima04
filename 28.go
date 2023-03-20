package main

import "fmt"

func main() {
	var num int
	var x int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	x = 1
	fmt.Println("Os divisores de", num, "são:")
	for num%x != 0 || num%x == 0 {
		x++
		if num%x == 0 {
			fmt.Println(x)
		}
	}

}
