// main.go

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var router *gin.Engine

func main() {

	r := gin.Default()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")

	// Handle Index
	r.GET("/", showIndexPage)
	// Handle GET requests at /article/view/some_article_id
	r.GET("/article/view/:article_id", getArticle)

	r.Run()

}
