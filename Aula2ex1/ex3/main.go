package main

import "fmt"

func calcWage(category string, hours float64) float64{
	var result float64
	switch category {
	case "C":
		result = 1000 * hours
	case "B":
		result = 1500 * hours
		if hours > 160 {
			result *= 1.20
		}
	case "A":
		result = 3000 * hours
		if hours > 160 {
			result *= 1.5
		}
	}
	return result
}

func main()  {
	fmt.Println(calcWage("A", 160.00))
	fmt.Println(calcWage("C", 180.00))
}

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func calculaSalario(categoria string, horas int) (float64, error) {
// 	if categoria == "C" {
// 		return float64(horas) * 1000, nil
// 	}
// 	if categoria == "B" {
// 		salario := float64(horas) * 1500
// 		if horas > 160 {
// 			return salario * 1.2, nil
// 		}
// 		return salario, nil
// 	}
// 	if categoria == "A" {
// 		salario := float64(horas) * 3000
// 		if horas > 160 {
// 			return salario * 1.5, nil
// 		}
// 		return salario, nil
// 	}
// 	return 0.0, errors.New("Categoria Inválida")
// }
// func main() {
// 	salario, _ := calculaSalario("C", 160)
// 	fmt.Printf("Salario 01: %.2f\n", salario)
// 	salario, _ = calculaSalario("B", 160)
// 	fmt.Printf("Salario 02: %.2f\n", salario)
// 	salario, _ = calculaSalario("A", 172)
// 	fmt.Printf("Salario 03: %.2f\n", salario)
// }
