package main

import (
	"sync"
	"runtime"
	"time"
	"log"
)

type st struct{
	A int32
	A1 int32
	A2 int32
	A3 int32
	A4 int32
	A5 int32
	A6 int32
	A7 int32
	A8 int32
	A9 int32
	A10 int32
	A11 int32
	A12 int32
	B int32
	Mutex sync.Mutex
}

func main(){
	runtime.GOMAXPROCS(50)
	var mp map[string]st
	mp = make(map[string]st)
	mp["k"] = st{A:-1,B:2,Mutex:sync.Mutex{}}

	//并发的修改map
	for i:= 1;i<=1000;i++{
		go func(ii int){
			v := mp["k"]
			v.A = int32(ii)
			v.A1 = int32(ii)
			v.A2 = int32(ii)
			v.A3 = int32(ii)
			v.A4 = int32(ii)
			v.A5 = int32(ii)
			v.A6 = int32(ii)
			v.A7 = int32(ii)
			v.A8 = int32(ii)
			v.A9 = int32(ii)
			v.A10 = int32(ii)
			//v.Mutex.Lock()
			mp["k"] = v
			//v.Mutex.Unlock()
		}(i)

		go func(jj int){
			v := mp["k"]
			if v.A ==0{
				log.Println("illegal happened!")
			}
		}(i)
	}

	//for j:=0;j<=100000;j++{
	//	go func(jj int){
	//		v := mp["k"]
	//		if v.A ==0{
	//			log.Println("illegal happened!")
	//		}
	//	}(j)
	//}

	time.Sleep(500 * time.Second)
}