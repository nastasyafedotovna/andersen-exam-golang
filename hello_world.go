package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func HelloWorld(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Exam ansersen"})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("index.html")

	router.GET("/", HelloWorld)
	router.Run()

}
