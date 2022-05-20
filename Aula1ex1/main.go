package main

import "fmt"

var nome string = "Renata"
var idade int = 32

var (
	temperatura = 27
	umidade = 90
	pressao = 40
)

func main()  {
	fmt.Println(nome)
	fmt.Println(idade)

	fmt.Println(temperatura, umidade, pressao)
}