package main

import (
	"fmt"
	"time"
	"runtime"
	"flag"
)

func main(){
	//注意：当设置P为1的时候，会进入无限for循环中；当设置P为2的时候，下面的100个数值能够正常打印，程序会结束
	//设置为2的时候，其他g有机会分配到一个空闲的p上面
	//设置为1的时候，空转for占据一个P，其他的g得不到运行的机会；如果是系统调用，则还是有机会的。
	//Q：单个p的时候，如果是一个g进入系统调用之后，阻塞了。此时其他的g有机会运行么？
	var procs int64
	flag.Int64Var(&procs,"P",1,"GOMAXPROCS的个数")
	flag.Parse()

	runtime.GOMAXPROCS(int(procs))
	go func(){
		for{}
	}()
	time.Sleep(time.Second)
	go func(){
		for i:=0;i<=100;i++{
			fmt.Println(i)
		}
	}()

	time.Sleep( 8 *time.Second)
}
