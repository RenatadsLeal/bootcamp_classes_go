package main

import (
	"errors"
	"fmt"
)

const (
	dog = "dog"
	cat = "cat"
	hamster = "hasmter"
	tarantula = "tarantula"
)

func dogFood(quantity int) int {
	return quantity * 10000
}

func catFood(quantity int) int {
	return quantity * 5000
}

func hamsterFood(quantity int) int {
	return quantity * 250
}

func tarantulaFood(quantity int) int {
	return quantity * 150	
}

func Animal(animalType string) (func(quantity int) int, error)  {
	if animalType == dog {
		return dogFood, nil
	}
	if animalType == cat {
		return catFood, nil
	}
	if animalType == hamster {
		return hamsterFood, nil
	}
	if animalType ==tarantula {
		return tarantulaFood, nil
	}
	return nil, errors.New("Invalid animal type!")
}

func main()  {

	animalDog, _ := Animal(dog)
	fmt.Printf("Quantidade necessária de alimento: %d gramas \n", animalDog(5))

	animalHamster, _ := Animal(hamster)
	fmt.Printf("Quantidade necessária de alimento: %d gramas \n", animalHamster(3))
}