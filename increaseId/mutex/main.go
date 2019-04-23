package main

import (
	"runtime"
	"log"
	"time"
	"sync"
)

//全局唯一id
var GId int64
var mtx sync.Mutex

func main(){
	runtime.GOMAXPROCS(5)
	for i:= 0;i<100;i++{
		go func(){
			num := GetId()
			log.Println(num)
		}()
	}
	time.Sleep(5 * time.Second)
}

func GetId()int64{
	mtx.Lock()
	GId = GId +1
	mtx.Unlock()
	return GId
}
