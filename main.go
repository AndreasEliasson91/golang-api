package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageView struct {
	Title  string
	Header string
}

var randomness *rand.Rand

func startApp(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", &PageView{Title: "test", Header: "Golang"})
}
