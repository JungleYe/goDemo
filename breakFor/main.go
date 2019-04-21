package main

import (
	"log"
	"time"
)

func main(){

	//核心提示：A "break" statement terminates execution of the innermost "for", "switch", or "select" statement within the same function.

	//普通跳出for循环测试
	for i:=0;i<100;i++{
		log.Printf("the i= %d",i)
		if i>= 10{
			break
		}
	}

	//定时器1：N长时间之后触发
	tmr := time.NewTimer(3 *time.Second)
	var cnt int32
Lp0:
	for{
		select{
		case <- tmr.C:
			log.Printf("时间到了：now = %s",time.Now().String())
			tmr.Reset(3 *time.Second) //重置tmr
			cnt = cnt+1
			if cnt >= 10{
				break Lp0//跳出for循环
				//此处如果直接使用break会无效果。不会退出for；
			}
		}
	}
}
