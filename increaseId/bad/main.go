package main

import (
	"runtime"
	"time"
	"sync"
	"fmt"
)

//全局唯一id
var GId int64
var mtx sync.Mutex

func main(){
	runtime.GOMAXPROCS(5)
	for i:= 0;i<1000;i++{
		go func(){
			num := GetId()
			fmt.Println(num)
		}()
	}
	time.Sleep(5 * time.Second)
}

func GetId()int64{
	GId = GId +1
	return GId
}
