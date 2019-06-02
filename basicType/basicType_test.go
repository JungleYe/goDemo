package basicType

import (
	"testing"
)

func TestInt(t *testing.T){
	var a int
	if a != 0{
		t.Errorf("未初始化的int值不等于0")
	}
	t.Logf("未初始化的int值等于0")
}

func TestStructPoint(t *testing.T){
	type st struct{
		name string
		age int32
	}
	var oneSt *st
	if oneSt == nil{
		t.Logf("未初始化的结构体指针等于nil")
	}

	//【panic】:未初始化的结构体指针，不能够访问其结构体成员
	//panic: runtime error: invalid memory address or nil pointer dereference
	//fmt.Print(oneSt.name)
}

func TestSlice(t *testing.T){
	type st struct{
		name string
		age int32
	}
	var one []st
	if one == nil{
		t.Logf("未初始化的slice等于nil")
	}

	//【panic】:未初始化的slice不能访问其下标地址
	//panic: runtime error: index out of range
	//fmt.Print(one[0])

	//【panic】:未初始化的slice不能访问其下标地址
	//panic: runtime error: index out of range
	//var two = make([]st,5)
	//fmt.Print(two[5])
}

func TestChannel(t *testing.T){
	var ch chan struct{}
	if ch == nil{
		t.Logf("未初始化的channel等于nil")
	}else{
		t.Logf("未初始化的channel不等于nil")
	}
	//【panic】:未初始化的channel会出现
	//fatal error: all goroutines are asleep - deadlock!
	//go func(){
	//	for {
	//		select{
	//		case v := <- ch:
	//			t.Logf("从通道中拿到了数据=%v",v)
	//		}
	//	}
	//}()
	//time.Sleep(time.Second)
	//ch <- struct{}{}
	//time.Sleep(time.Second)
}

