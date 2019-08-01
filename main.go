package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = io.MultiWriter(os.Stdout)
	fmt.Print("初始化")

}

func main() {
	fmt.Print("主函数")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
