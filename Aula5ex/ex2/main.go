package main

import (
	"errors"
	"fmt"
	"log"
)

func checkWage(wage float64) {
	if wage < 15.000 {
		fmt.Println(errors.New("O salário digitado não está dentro do valor mínimo"))
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}

}

/////////////////////////////////////////

func checkWage2(wage float64) (string, error) {
	if wage < 15.000 {
		return "", errors.New("O salário digitado não está dentro do valor mínimo")
	}

	return "Necessário pagamento de imposto", nil

}

func main() {

	salario := 5.000

	checkWage(salario)

	////////////////////////////////////

	msg, err := checkWage2(salario)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(msg)

}
