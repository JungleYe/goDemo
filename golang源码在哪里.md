```golang
package main

import (
	"math"
	"fmt"
	"time"
)

func main(){
	//查看源码的方式一：查看math.Asin源码的方法：
	//标准库中的方法只有：func Asin(x float64) float64
	//同级目录中存在对各种体系结构专用的汇编文件asin_386.s  asin_adm64.s 此处就是其真实所在
	val := math.Asin(0.2)

	//查看源码的方式二：查看fmt.Println源码的方法：
	//标准库中就有其完整的实现，比较好找
	fmt.Println(val)

	//查看源码的方式三：查看time.Sleep的源码的方法：
	//标准库中的方法为：func Sleep(d Duration) 没有具体的实现
	//通过 `grep linkname /usr/local/go/src/runtime/*.go | grep time.Sleep` 找到其所在地址为：
	// /usr/local/go/src/runtime/time.go://go:linkname timeSleep time.Sleep
	//到 /usr/local/go/src/runtime/time.go去找对用的实现就可以了
	time.Sleep(time.Second)
}
```
