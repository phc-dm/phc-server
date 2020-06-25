package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.Static("/blog", "./blog/public")

	r.GET("/", func(c *gin.Context) {
		r.SetHTMLTemplate(template.Must(template.ParseFiles("./views/base.html", "./views/home.html")))
		c.HTML(http.StatusOK, "base", gin.H{})
	})
	r.GET("/utenti", func(c *gin.Context) {
		r.SetHTMLTemplate(template.Must(template.ParseFiles("./views/base.html", "./views/utenti.html")))
		c.HTML(http.StatusOK, "base", gin.H{})
	})

	r.Run(":8000")
}
