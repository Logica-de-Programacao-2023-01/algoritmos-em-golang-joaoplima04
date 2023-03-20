package main

import "fmt"

func main() {
	var contagem int
	for contagem != 30 {
		contagem++
		if contagem%3 == 0 {
			fmt.Println(contagem)
		}
	}

}
