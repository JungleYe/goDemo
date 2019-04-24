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

var Sum int64
var sumLock sync.Mutex
var wg sync.WaitGroup

func main(){
	runtime.GOMAXPROCS(5)
	for i:= 0;i<1234567;i++{
		wg.Add(1)
		go func(){
			num := GetId()
			addSum(num)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("the mutex total = %d",Sum) //762078472657
	time.Sleep(5 * time.Second)
}

//FATAL ERROR :并发情况下直接return GId非常危险！！！在return GId之前很可能还是其他goroutine执行了加一动作了。
func GetIdInFatalUsage()int64{
	mtx.Lock()
	GId = GId +1
	mtx.Unlock()
	return GId
}

func GetId()int64{
	mtx.Lock()
	GId = GId +1
	n := GId
	mtx.Unlock()
	return n
}

func addSum(ad int64){
	sumLock.Lock()
	Sum = Sum + ad
	sumLock.Unlock()
}