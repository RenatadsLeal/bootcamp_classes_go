package main

import "fmt"

func main() {

	contador := 0

	var fruit string = "banana"
	for _, letter := range fruit {

		contador++
		fmt.Printf("%c\n", letter)
	}

	fmt.Printf("%s tem %d letras", fruit, contador)
}
