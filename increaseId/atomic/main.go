package main

import (
	"sync/atomic"
	"runtime"
	"log"
	"time"
)

//全局唯一id
var GId int64

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
	var c int64	//current
	var n int64	//next
	for{
		c = atomic.LoadInt64(&GId)
		n = c+1
		if atomic.CompareAndSwapInt64(&GId,c,n){
			return n
		}
	}
}
