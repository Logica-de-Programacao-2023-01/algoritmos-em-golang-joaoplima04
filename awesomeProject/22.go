package main

import "fmt"

func main() {
	var contagem int
	for contagem != 20 {
		contagem++
		if contagem%2 == 0 {
			fmt.Println(contagem)
		}
	}
}
