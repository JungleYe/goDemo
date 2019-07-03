package main

import (
	"time"
	"fmt"
)

type st struct{
	name string
	age int32
}
var ch  chan *st
func main(){

	//知识点：没有经过make的channel，就等着被锁死吧；windows下会有这个提示：fatal error: all goroutines are asleep - deadlock!
	//上面错误提示，是没有其余goroutines，这个channel不可能会收到消息了，触发的panic
	ch = make(chan *st)
	go Dispatch()

	time.Sleep( time.Second)
	var oneSt = &st{"jungle",29}
	ch <- oneSt
	time.Sleep(2 * time.Second)
	fmt.Println("the end")
}

func  Dispatch() error {
	fmt.Println("开启了Dispatch")
	for {
		select {
		case req := <- ch:
			fmt.Println("可以从channel中拿到数据了！")
			fmt.Println(req)
		}
	}
	return nil
}
