package main

import (
	"sync/atomic"
	"runtime"
	"time"
	"sync"
	"fmt"
)

//全局唯一id
var GId int64

//计算总数，核对结果是否正常
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
	fmt.Printf("the atomic total = %d",Sum) //762078456028
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

func addSum(ad int64){
	sumLock.Lock()
	Sum = Sum + ad
	sumLock.Unlock()
}
