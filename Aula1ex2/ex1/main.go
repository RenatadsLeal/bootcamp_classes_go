package main

import "fmt"

func main() {

	contador := 0

	var fruit string = "banana"
	for _, letter := range fruit {

		contador++
		fmt.Printf("%c\n", letter)
	}

	fmt.Printf("%s tem %d letras \n", fruit, contador)


	//////////////////////////////////////////////////

	bootcamp := "Bootcamp GO W1"

	fmt.Println("tamanho:", len(bootcamp))

	for _, r := range bootcamp {
		fmt.Println(string(r))
	}


}
