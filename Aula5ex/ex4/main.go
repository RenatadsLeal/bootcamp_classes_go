package main

import (
	"errors"
	"fmt"
)

// type myError struct {
// 	msg string
// }

// func (e myError) Error() string {
// 	return fmt.Sprintf("erro: %s", e.msg)
	
// }

func calcWage(hours int, price float64) (float64, error) {
	wage := float64(hours) * price

	if wage >= 20.000 {
		wage = wage - wage/10
	}

	if hours < 80 {
		return 0, errors.New("horas inválidas")
	}

	return wage, nil
	
}

func main()  {

	hours := 60
	price := 15.00

	msg, err := calcWage(hours, price)
	if err != nil {
		err2 := fmt.Errorf("tipo de erro: %w", err)
		fmt.Println(err2)
		fmt.Println(errors.Unwrap(err2))
	} else {
		fmt.Println(msg)
	}

}