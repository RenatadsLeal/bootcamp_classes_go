package main

import (
	"fmt"

	"github.com/RenatadsLeal/bootcamp_classes_go/Aula1/pkg/calc"
)

var nome2 string = "Renata"

/**
const nome3 string = "Renata"
**/
var ligarLuzes bool = false

const (
	pessoa1           = "Helena"
	pessoa2           = "Carol"
	quantidadePessoas = 2
)

const dev = "asa"

var primeiroNome, sobrenome = "Renata", "Leal"

func main() {
	fmt.Println("Hello World")

	// var nome string
	// nome = "Renata"

	nome := "Renata" // atribuição curta não funciona fora de funções

	fmt.Println(nome)
	fmt.Println(nome2)
	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	fmt.Println(ligarLuzes)
	fmt.Println(sobrenome)

	fmt.Println(calc.Sum(5, 5))
	fmt.Println(calc.PI)

	var a [2]string // array

	a[0] = "Hello"
	a[1] = "World"

	b := make([]string, 5) // slice

	var c []string  //slice

	c = append(c, "elemento") // slice tem append

	fmt.Println(c)

	var d [3]string
	fmt.Println(d)
	fmt.Println(len(d))

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(a[0])

	//map

	var students = make(map[string]interface{})

	students["Maria"] = 20
	students["Celso"] = "opa"

	fmt.Println(len(students))
	fmt.Println(students)

	delete(students, "Celso")
	fmt.Println(students)

	//for

	frutas := []string{"maçã", "banana", "pêra"}
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}

	for key, fruta := range frutas {
		fmt.Println(key, fruta)
	}

	// loop while
	sum := 1

	for sum < 10 {
		sum += sum
	}

	fmt.Println(sum)

	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i, "is odd")
	}


}
