package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type Product struct {
    Id         int64   `json:"id"`
    Nome       string  `json:"nome"`
    Tipo       string  `json:"tipo"`
    Quantidade int     `json:"quantidade"`
    Preco      float64 `json:"preco"`
}

type Products []Product

var lastID int64
var products Products

func listProducts(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": &products})
}

func saveProduct(c *gin.Context) {
    var produto Product

    if err := c.ShouldBindJSON(&produto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    lastID++
    produto.Id = lastID

    products = append(products, produto)

    c.JSON(200, gin.H{
        "data": produto,
    })
}

func main() {
    r := gin.Default()

    group := r.Group("/produtos")
    {
        group.GET("/", listProducts)
        group.POST("/", saveProduct)
    }

    r.Run()
}