package main

import "fmt"

var employees = map[string]int{"Benjamim": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

func main()  {
	fmt.Println("Bejamim tem", employees["Benjamim"], "anos" )

	fmt.Println(employees)

	employees["Frederico"] = 25
	fmt.Println(employees)

	delete(employees, "Pedro")
	fmt.Println(employees)
	
}