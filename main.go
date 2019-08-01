package main

import (
	"OnlineGobangBattle/config"
	"OnlineGobangBattle/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = io.MultiWriter(os.Stdout)
	fmt.Print("初始化")
	config.Loadconf()
	service.ConnectRedis()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()

}
