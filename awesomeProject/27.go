package main

import "fmt"

func main() {
	var contagem int
	for contagem != 100 {
		contagem++
		if contagem%3 == 0 {
			fmt.Print("Fizz")
		}
		if contagem%5 == 0 {
			fmt.Print("Buzz")
		}
		if contagem%3 == 0 && contagem%5 == 0 {
			fmt.Print("FizzBuzz")
		}
		fmt.Println(contagem)

	}

}
