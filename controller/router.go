package controller

import (
	"OnlineGobangBattle/model"
	"OnlineGobangBattle/service"
	"OnlineGobangBattle/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"strings"
)

func MapRouters() *gin.Engine {
	ret := gin.New()
	ret.Static("/static", "./StaticRes")
	store := cookie.NewStore([]byte(""))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		Secure:   strings.HasPrefix("http://192.168.18.113:8080", "http"),
		HttpOnly: true,
	})
	ret.Use(sessions.Sessions("mysession", store))
	ret.LoadHTMLGlob("./template/*")
	ret.GET("/ws", func(c *gin.Context) {
		// change the reqest to websocket model
		conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
		if error != nil {
			http.NotFound(c.Writer, c.Request)
			return
		}
		session := sessions.Default(c)
		value := session.Get("uid")
		// websocket connect
		client := &service.Client{ID: value.(string), Socket: conn, Send: make(chan []byte)}
		service.Manager.Register <- client
		go client.Read()
		go client.Write()
	})

	ret.GET("/setuser", func(c *gin.Context) {
		uid := uuid.NewV4().String()
		name := utils.GetFullName()
		session := sessions.Default(c)
		value := session.Get("uid")
		if value != nil {

			Conn := service.RedisClient.Get()
			results, err := redis.Values(Conn.Do("lrange", "Listofusers", 0, -1))
			if err != nil {
				log.Println("redis add key   failure", err)
			}
			var rslist []string
			for _, v := range results {
				rslist = append(rslist, string(v.([]byte)))
			}
			var hasvalue bool = false
			for _, v := range rslist {
				if v == value.(string) {
					hasvalue = true
					break
				}
			}
			if hasvalue == false {
				_, err := Conn.Do("RPUSH", "Listofusers", value.(string))
				if err != nil {
					log.Println("redis add key   failure", err)
				}
				Conn.Close()
			}
			var usereentity model.User
			json.Unmarshal([]byte(value.(string)), &usereentity)
			c.JSON(200, gin.H{
				"uid":  usereentity.Uuid,
				"name": usereentity.Name,
			})

		} else {
			j, errj := json.Marshal(model.User{
				Name: name,
				Uuid: uid,
			})
			if errj != nil {
				log.Println("json to data failure ", errj)
			}
			session.Set("uid", string(j))
			session.Options(sessions.Options{
				Domain: "192.168.18.113",
			})
			session.Save()
			//写入  redis 记录
			Conn := service.RedisClient.Get()
			_, err := Conn.Do("RPUSH", "Listofusers", j)
			if err != nil {
				log.Println("redis add key   failure", err)
			}
			Conn.Close()
			c.JSON(200, gin.H{
				"uid":  uid,
				"name": name,
			})
		}

	})

	ret.GET("/user", func(c *gin.Context) {
		session := sessions.Default(c)
		value := session.Get("uid")
		c.JSON(200, gin.H{
			"uid": value,
		})
	})

	ret.GET("/online", func(c *gin.Context) {
		session := sessions.Default(c)
		value := session.Get("uid")
		if value == nil {
			c.JSON(200, gin.H{
				"online": "[]",
			})
		} else {
			Conn := service.RedisClient.Get()
			results, err := redis.Values(Conn.Do("lrange", "Listofusers", 0, -1))
			if err != nil {
				log.Println("redis add key   failure", err)
			}
			var rslist []string
			for _, v := range results {
				rslist = append(rslist, string(v.([]byte)))
			}

			Conn.Close()
			c.JSON(200, gin.H{
				"online": rslist,
			})
		}
	})

	ret.GET("/Createaroom", func(c *gin.Context) {
		session := sessions.Default(c)
		value := session.Get("uid")
		if value == nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"state": false,
			})
		} else {

			name := c.Query("name")
			fmt.Print(name)
			fmt.Print("获取参数")

			c.JSON(200, gin.H{
				"state": true,
				"info":  name,
			})

		}

	})

	ret.GET("/getroomlist", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "War",
		})
	})

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
