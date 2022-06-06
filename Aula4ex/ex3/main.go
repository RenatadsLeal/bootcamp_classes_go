package main

import "fmt"

type produto struct {
	nome string
	preco float64
	quantidade int
}

type servico struct {
	nome string
	preco float64
	minuto int
}

type manutencao struct {
	nome string
	preco float64
}

func somarProdutos(produtos []produto, c chan float64) float64 {
	var total float64
	for _, prod := range produtos {
		total += prod.preco * float64(prod.quantidade)
	}
	fmt.Println("total produto: ", total)
	c <- total
	return total
	
}

func somarServicos(servicos []servico, c chan float64) float64 {
	var totalHoras int
	var totalPreco float64
	var total float64
	for _, serv := range servicos {
		totalHoras += serv.minuto
		totalPreco += serv.preco

	}
	mediaHoras := totalHoras / (len(servicos))
	total = totalPreco * float64(mediaHoras)
	fmt.Println("total servico: ", totalPreco * float64(mediaHoras))
	c <- total
	return total
	
}

func somarManutencao(manut []manutencao, c chan float64) float64 {
	var total float64
	for _, a := range manut {
		total += a.preco
	}
	fmt.Println("total manutencao: ", total)
	c <- total
	return total
	
}

func main()  {
	c1 := make(chan float64)

	produto1 := produto{nome: "relógio", preco: 50, quantidade: 2}
	produto2 := produto{nome: "caderno", preco: 10.50, quantidade: 1}

	meusProdutos := []produto{produto1, produto2}

	// fmt.Println(somarProdutos(meusProdutos))

	go somarProdutos(meusProdutos, c1)

	c2 := make(chan float64)

	servico1 := servico{nome: "instalacao de ar", preco: 8, minuto: 90}
	servico2 := servico{nome: "reparo de ar", preco: 2, minuto: 60}

	meusServicos := []servico{servico1, servico2}

	go somarServicos(meusServicos, c2)

	// fmt.Println(somarServicos(meusServicos))

	c3 := make(chan float64)

	manutencao1 := manutencao{nome: "conserto de relógio", preco: 100}
	manutencao2 := manutencao{nome: "troca de bateria", preco: 40}
	totalManutencoes := []manutencao{manutencao1, manutencao2}

	go somarManutencao(totalManutencoes, c3)

	// fmt.Println(somarManutencao(totalManutencoes))

	somaTotal := <-c1 + <-c2 + <-c3

	fmt.Println("soma com channel:", somaTotal)
	
}