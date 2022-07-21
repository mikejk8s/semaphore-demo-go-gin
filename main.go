// main.go

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var router *gin.Engine

func main() {

	router := gin.Default()
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.LoadHTMLGlob("templates/*")

	// Handle Index
	router.GET("/", showIndexPage)
	// Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", getArticle)

	router.Run()

}
