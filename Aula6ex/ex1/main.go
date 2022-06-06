package main

import "os"

func main()  {
	_,err := os.Open("customers.txt")
	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}
}