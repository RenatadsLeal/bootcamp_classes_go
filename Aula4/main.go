package main

import "fmt"

func main() {

	var p *int        // ponteiro nill
	var p2 = new(int) //ponteiro

	// * para criar variável do tipo pointeiro
	// * para visualizar o valor da posição da memória
	// & para atribuir um valor ao ponteiro e qaundo passa o valor

	fmt.Println("O endereço da variável p é: ", p)
	fmt.Println("O endereço da variável p é: ", p2)

	// quando recebe o valord e ponteiro precisa o usar o &

	var v int = 19
	var t *int = &v

	fmt.Println("O endereço da variável t é: ", t) //local memória
	fmt.Println("O valor da variável t é: ", *t) //valor variável

}
