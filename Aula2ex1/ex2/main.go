package main

import "fmt"


func finalGrade(grades ...float64)  float64{
	var result float64
	for _,grade := range grades {
		result += grade
	}
	return result / float64(len(grades))
}

func main()  {
	fmt.Println(finalGrade(10, 9.5, 9.5, 10))	
}