package main

import (
	"runtime"
	"time"
	"sync"
	"fmt"
)

//全局唯一id
var GId int64

var Sum int64
var sumLock sync.Mutex
var wg sync.WaitGroup

func main(){
	runtime.GOMAXPROCS(5)
	for i:= 0;i<1234567;i++{
		num := GetId()
		addSum(num)
	}
	fmt.Printf("the oneGoutine total = %d",Sum)	// 762078456028
	time.Sleep(5 * time.Second)
}

func GetId()int64{
	GId = GId +1
	return GId
}

func addSum(ad int64){
	sumLock.Lock()
	Sum = Sum + ad
	sumLock.Unlock()
}
