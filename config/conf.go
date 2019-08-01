package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

var Conf *Configuration

type Configuration struct {
	RedisMaxIdle        int
	RedisMaxActive      int
	RedisMaxIdleTimeout int
	Port                string
	RedisHost           string
	Redispwd            string
	RedisDb             int
}

func Loadconf() {
	confPath := flag.String("conf", "config.json", "path of config.json")
	bytes, err := ioutil.ReadFile(*confPath)
	if nil != err {
		log.Print("loads configuration file [" + *confPath + "] failed: " + err.Error())
	}
	Conf = &Configuration{}
	if err = json.Unmarshal(bytes, Conf); nil != err {
		log.Print("parses [pipe.json] failed: ", err)
	}
	fmt.Print("加载配置文件成功")
}
