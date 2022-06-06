package main

import (
	"fmt"
	"log"
)

func logging() {
	log.Println("Logando")

}

//////////////////////////

func soma(a, b float32) float32 { // se forem do mesmo tipo seguido pode especificar só no final
	return a + b
}

///////////////////////////

func somaNums(values ...float64) float64 { //ellipsis ou função variádica
	var result float64
	for _, value := range values {
		result += value
	}
	return result
}

///////////////////////////

func foo(str string, fn func(value string) string) string {
	return fn(str)
}

func hello(value string) string {
	return "Hello, " + value
}

///////////////////////////

type Pessoa struct {
	Nome      string
	Gênero    string
	Idade     int
	Profissão string
	Peso      float64
}

type Pessoas []Pessoa

///////////////////////////

type Pessoa2 struct {
	PrimeiroNome string `json:"primeiro_nome" bd:"primeiro_nomee"`
	Sobrenome    string `jason:"sobrenome"`
	Telefone     string
	Endereco     string `json:"endereco"`
}

func main() {
	logging()
	fmt.Println("Logando")

	fmt.Println(soma(5, 5))

	minhaSoma := soma(20, 20)
	fmt.Println(minhaSoma)

	fmt.Println(somaNums(5, 100, 50, 50, 100))

	fmt.Println(foo("World", hello))

	renata := Pessoa{
		Nome:      "renata",
		Gênero:    "Feminino",
		Idade:     32,
		Profissão: "Software Developer",
		Peso:      60,
	}

	carol := Pessoa{}
	var helena Pessoa

	carol.Nome = "Carol"
	helena.Idade = 31

	renata.Peso = 59

	fmt.Println(renata.Profissão)
	fmt.Println("Idade: ", carol.Idade)

	// fmt.Println(string(meuJSON))
}
