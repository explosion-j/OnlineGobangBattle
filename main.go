package main

import (
	"OnlineGobangBattle/config"
	"OnlineGobangBattle/controller"
	"OnlineGobangBattle/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = io.MultiWriter(os.Stdout)
	fmt.Print("初始化")
	config.Loadconf()
	service.ConnectRedis()
	//
}

func main() {
	go service.Manager.Start()
	router := controller.MapRouters()
	server := &http.Server{
		Addr:    "0.0.0.0:" + config.Conf.Port,
		Handler: router,
	}
	handleSignal(server)
	if err := server.ListenAndServe(); nil != err {
		log.Fatalf("listen and serve failed: " + err.Error())
	}

}

func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		s := <-c
		log.Panicf("got signal [%s], exiting gobang now", s)
		if err := server.Close(); nil != err {
			log.Printf("server close failed" + err.Error())
		}
		log.Print(" exited")
		os.Exit(0)
	}()
}
