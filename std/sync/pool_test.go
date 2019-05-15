package main

import (
	"sync"
	"testing"
)

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

//func Benchmark_WithPool(b *testing.B){
//	for i:=0;i<=b.N;i++{
//		obj := bytePool.Get().(*[]byte)
//		v := obj
//		_ = v
//		bytePool.Put(obj)
//	}
//}