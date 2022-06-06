package main

import (
	"fmt"
	"time"
)

type aluno struct {
	Nome string
	Sobrenome string
	RG string
	DataAdmissao time.Time 
}

func (a aluno) Detalhe() {
	fmt.Printf("Nome completo do aluno: %s %s\n", a.Nome, a.Sobrenome)
	fmt.Println("RG:", a.RG)
	fmt.Println("Data de admissão:", a.DataAdmissao)
}

func main()  {

	aluno1 := aluno {
		Nome: "Alan",
		Sobrenome: "Souza",
		RG: "075463",
		DataAdmissao: time.Date(2019, 05, 12, 0, 0, 0, 0, time.UTC),
	}
	
	aluno1.Detalhe()
}