package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type product struct {
	Id         int
	Nome       string
	Cor        string
	Preco      float64
	Estoque    int
	Codigo     string
	Publicacao bool
}

// func GetAll(w http.ResponseWriter, req *http.Request) {
// 	product1 := product{Id: 37454, Nome: "celular", Cor: "preto", Preco: 3.550, Estoque: 250, Codigo: "hy43e", Publicacao: true}
// 	product2 := product{Id: 33354, Nome: "computador", Cor: "cinza", Preco: 5.050, Estoque: 210, Codigo: "dhdyt78", Publicacao: true}
// 	prods := []product{product1, product2}
// 	fmt.Fprintf(w, prods)
// }

func main() {

	router := gin.Default()

	router.GET("/ola", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Olá, Renata",
		})
	})

	// http.HandleFunc("/products", GetAll)
	// http.ListenAndServe(":8080", nil)

	router.Run()

}
