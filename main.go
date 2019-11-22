package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("tem/*")
	r.GET("/", GetBook)
	r.Run()
}
