package main

import (
	"time"
	"log"
)

func main(){
	useTimeStamp()
}

func useTimeStamp(){
	//以int64的方式存储时间戳
	var timeStamp int64
	timeStamp = time.Now().Unix()
	log.Printf("当前时间戳为 %d",timeStamp)

	time.Sleep(5 * time.Second)

	//打印出时间
	log.Printf("以自定义的时间格式打印：%s",time.Unix(timeStamp,0).Format("2006-01-02 15:04:05"))

	//查看时间间隔duration
	dr := time.Since(time.Unix(timeStamp,0))
	//打印时间间隔
	log.Printf("the duration is %s",dr.String())

	/**
		ticker只要定义完成，从此刻开始计时，不需要任何其他的操作，每间隔固定时间都会触发；
		timer定时器，是到了固定事件后会执行一次
		如果timer定时器要实现ticker的效果，使用func (t *Timer) Reset(d Duration) bool来重置定时器
	 */
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
			}
		}
	}

	//time.Ticker
	ticker := time.NewTicker(5 * time.Second)
	for{
		select{
		case <- ticker.C:
			log.Printf("ticker时间到:%s",time.Now().String())
		}
	}
}
