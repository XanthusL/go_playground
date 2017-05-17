package main

import (
	"net/http"

	"strconv"

	"src/github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("template/*")
	initializeRouters()
	router.Run()
}

func initializeRouters() {
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
}
func showIndexPage(c *gin.Context) {
	a := getAllArticles()
	render(c, gin.H{
		"title":   "Home Page",
		"payload": a,
	}, "index.html")
}
func getArticle(c *gin.Context) {
	aid := c.Param("article_id")
	i, _ := strconv.ParseInt(aid, 10, 32)
	a := getArticleById(int(i))
	if a == nil {
		c.String(http.StatusNotFound, "asdf", nil)
		return
	}
	render(c, gin.H{
		"title":   a.Title,
		"payload": a,
	}, "article.html")

}
func render(c *gin.Context, data gin.H, tmpl string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, tmpl, data)
	}
}
