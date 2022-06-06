package main

import "fmt"

var employees = map[string]int{
	"Benjamim": 20,
	"Manuel":   26,
	"Brenda":   19,
	"Dario":    44,
	"Pedro":    30,
}

func main() {
	fmt.Println("Bejamim tem", employees["Benjamim"], "anos")

	fmt.Println(employees)

	c := 0
	for _, v := range employees {
		if v > 21 {
			c++
		}
	}

	fmt.Printf("Funcionários com mais de 21 anos: %d \n", c)

	employees["Frederico"] = 25
	fmt.Println(employees)

	delete(employees, "Pedro")
	fmt.Println(employees)

}
