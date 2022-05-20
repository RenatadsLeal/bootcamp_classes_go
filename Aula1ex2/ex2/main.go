package main

import "fmt"

func concedeEmprestimo(idade int, empregado bool, anoAtividade int, salario float32)  {

	if idade < 22 || !empregado|| anoAtividade < 1 {
		fmt.Println("Sinto muito, empréstimo negado")
	} else {
		if salario > 100.000 {
			fmt.Println("Empréstimo concedido sem juros!")
		} else {
			fmt.Println("Empréstimo concedido com juros!")
		}
	}
	
}

func main()  {
	concedeEmprestimo(31, true, 2, 15.000)
	
}