package service

import (
	"github.com/gomodule/redigo/redis"
)

var redisClient *redis.Pool
var MaxIdled int = 1024

func ConnectRedis() {
	//conredis,err:= redis.Dial("tcp","127.0.0.1:6379")
	//if err!=nil{
	//	log.Println("redis connection  failure",err)
	//	return
	//}
	//maxIdle:=MaxIdled
	//if v,ok:=

}
