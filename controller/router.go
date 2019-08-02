package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MapRouters() *gin.Engine {
	ret := gin.New()
	ret.Static("/static", "./StaticRes")
	ret.LoadHTMLGlob("./template/*")
	ret.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "War",
		})
	})
	ret.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return ret
}
