package main

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "net/http"
  "strconv"

  "github.com/gin-gonic/gin"
)

type Products struct {
  Products []Product `json:"products"`
}

type Product struct {
  Id      int64   `json:"id"`
  Nome    string  `json:"nome"`
  Cor     string  `json:"cor"`
  Preco   float64 `json:"preco"`
  Estoque bool    `json:"estoque"`
  Codigo  string  `json:"codigo"`
}

func main() {
  r := gin.Default()

  r.GET("/hello", func(c *gin.Context) { // query ?name="Fulano" ou &lastname="Ciclano"
    name := c.Query("name")

    if name == "" {
      c.JSON(http.StatusOK, gin.H{
        "message": "Anonymous",
      })

      return
    }

    c.JSON(http.StatusOK, gin.H{
      "message": name,
    })
  })

  r.GET("/hello/:id", func(c *gin.Context) { // param /hello/:id
    id := c.Param("id")

    parsedID, err := strconv.Atoi(id)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": "erro interno, tente novamente",
      })

      log.Println(err.Error())

      return
    }

    if parsedID != 10 {
      c.JSON(http.StatusNotFound, gin.H{
        "error": "id não encontrado",
      })

      log.Println("id não encontrado")

      return
    }

    c.JSON(http.StatusOK, gin.H{
      "message": id,
    })
  })

  // Servindo arquivos usando a função File
  // r.GET("/products", func(c *gin.Context) {
  //  c.File("./products.json")
  // })

  // De forma manual
  r.GET("/products", func(c *gin.Context) {
    content, err := ioutil.ReadFile("./products.json")
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }

    var p Products

    if err := json.Unmarshal(content, &p); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }

    c.JSON(http.StatusOK, &p)
  })

  r.Run()
}