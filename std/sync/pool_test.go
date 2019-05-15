package main

import (
	"sync"
	"testing"
)

//重要：Len切换成65535和65536时的结果完全不一样！主要涉及到垃圾回收
//测试方式：go test -bench=".*" -cpuprofile=cpu.profile -benchmem
//
//测试结果：65536时
//pkg: goDemo/std/sync
//Benchmark_WithOutPool-4            50000             21527 ns/op           65537 B/op          1 allocs/op
//Benchmark_WithPool-4            30000000                36.9 ns/op             0 B/op          0 allocs/op
//
//pprof含义说明
//flat：给定函数上运行耗时
//flat%：同上的 CPU 运行耗时总比例
//sum%：给定函数累积使用 CPU 总比例 == 前一个sum+自己的flat占比
//cum：当前函数加上它之上的调用运行总耗时；cum代表的是该函数自身代码+所有调用的函数的执行时长；
//cum：我们以下面的例子来说明，函数b由三部分组成：调用函数c、自己直接处理一些事情、调用函数d，其中调用函数c耗时1秒，自己直接处理事情耗时3秒，调用函数d耗时2秒，那么函数b的flat耗时就是3秒，cum耗时就是6秒。
//cum%：同上的 CPU 运行耗时总比例
//
//pprof分析结果：
//Showing top 10 nodes out of 91
//      flat  flat%   sum%        cum   cum%
//     800ms 24.39% 24.39%      800ms 24.39%  runtime.memclrNoHeapPointers
//     290ms  8.84% 33.23%      640ms 19.51%  sync.(*Pool).Put
//     270ms  8.23% 41.46%      430ms 13.11%  runtime.scanobject
//     270ms  8.23% 49.70%      520ms 15.85%  sync.(*Pool).pin
//     170ms  5.18% 54.88%      420ms 12.80%  sync.(*Pool).Get
//     130ms  3.96% 58.84%      130ms  3.96%  runtime.procPin (inline)
//      90ms  2.74% 61.59%     1150ms 35.06%  goDemo/std/sync.Benchmark_WithPool
//      70ms  2.13% 63.72%       70ms  2.13%  sync.indexLocal (inline)
//      60ms  1.83% 65.55%      120ms  3.66%  runtime.findObject
//      60ms  1.83% 67.38%       60ms  1.83%  runtime.procUnpin (inline)

//分析：虽然看起来WithPool的cum比较高，但是，由上述benchmarck可知，这是其执行3000万次的结果
//TODO：两种方式单独运行，然后对比op和cpu消耗对比

//测试结果65535时
//TODO
const LEN int32 =  65536
var (
	bytePool = sync.Pool{
		New: func()interface{}{
			b := make([]byte,LEN)
			return &b
		},
	}
)

func Benchmark_WithOutPool(b *testing.B){
	for i:=0;i<=b.N;i++{
		obj := make([]byte,LEN)
		_ = obj
	}
}

func Benchmark_WithPool(b *testing.B){
	for i:=0;i<=b.N;i++{
		obj := bytePool.Get().(*[]byte)
		v := obj
		_ = v
		bytePool.Put(obj)
	}
}