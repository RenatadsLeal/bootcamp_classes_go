package main

import "fmt"

type myError struct {
	msg string
}

func (e *myError) Error() string {
	return fmt.Sprintf("erro: %s", e.msg)
}

func checkWage(wage float64) (string, error) {
	if wage < 15.000 {
		return "", &myError{"O salário digitado não está dentro do valor mínimo"}
	}
	return "Necessário pagamento de imposto", nil
	
}

func main()  {
	salario := 5.000

	fmt.Println(checkWage(salario))
	
}