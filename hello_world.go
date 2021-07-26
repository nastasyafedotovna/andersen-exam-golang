package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Exam ansersen"})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("index.html")

	router.GET("/", HelloWorld)
	err := router.Run()
	if err != nil {
		fmt.Print("fk u linter hehe")
	}

}
