package main

import (
	"fmt"
	"log"
)

func checkWage(wage float64) (string, error) {
	if wage < 15.000 {
		return "", fmt.Errorf("erro: o mínimo tributável é 15.000 e o salário informado é: %f", wage)
	}

	return "Necessário pagamento de imposto", nil

}

func main() {

	wage := 5.000

	msg, err := checkWage(wage)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(msg)

}
