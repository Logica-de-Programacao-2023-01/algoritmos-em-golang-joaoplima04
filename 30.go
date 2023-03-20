package main

import "fmt"

func main() {
	var listanum, num, max int
	fmt.Println("Digite vários numeros os separando com espaço. Quando terminar digite 0")
	fmt.Scan(&listanum)
	max = 0
	for listanum != 0 {
		fmt.Scan(&num)
		for num > max {
			num += max
			if num == 0 {
				break
				fmt.Println(max)
			}
		}
	}
}
