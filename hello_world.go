package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	
	//router.StaticFile("index.html", "index.html")
	router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{"title": "Exam ansersen"})
	})
	router.Run()

}
