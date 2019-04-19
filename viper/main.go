package main

import (
	"flag"
	"log"
	"github.com/spf13/viper"
	"fmt"
)

func main(){
	var configFile *string
	configFile = flag.String("c","./init.yaml","服务配置文件")
	flag.Parse()
	log.Println("configFile = %s",*configFile)

	viper.SetConfigFile(*configFile)
	err := viper.ReadInConfig()
	if err != nil{
		log.Printf("read configFile err:%v",err)
	}

	env := viper.GetString("Env")
	zkHost := viper.GetString(fmt.Sprintf("%s.zkHost",env))

	log.Printf("the zkHost = %s",zkHost)
}
