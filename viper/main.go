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

	//打印结果：2019/04/20 09:12:03 the zkHost = 127.0.0.1
	env := viper.GetString("Env")
	zkHost := viper.GetString(fmt.Sprintf("%s.zkHost",env))
	log.Printf("the zkHost = %s",zkHost)

	//打印结果：2019/04/20 09:12:03 the ips = [192.168.0.1 192.168.0.2 192.168.0.3]
	ips := viper.GetStringSlice("ip")
	log.Printf("the ips = %v",ips)
}
