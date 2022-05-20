package main

// ex1

import "fmt"

func countLetters(word string) {
	for _, letter := range word {
		fmt.Printf("%b\n", letter)
	}
}

func main() {

	// var fruit string = "apple"

	// for i, letter := range fruit {
	// 	fmt.Println(i, letter)
	// }

	// countLetters("banana")

	contador := 0

	var fruit string = "banana"
	for _, letter := range fruit {

		contador++
		fmt.Printf("%c\n", letter)
	}
	fmt.Println(contador)

}
